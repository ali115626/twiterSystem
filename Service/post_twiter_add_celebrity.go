package Service

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
	Dao "twiter/tweetSystem/Dao/mysql"
	DaoRedis "twiter/tweetSystem/Dao/redis"
	"twiter/tweetSystem/utitl"
)

//todo 测试一个明星和非明星


//

func PostTweetAddCelebrity(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	userId := requestMap["userId"][0]
	content := requestMap["content"][0]
	createAt:=time.Now()
	lastRowId,err :=Dao.InsertTweet(userId,content,createAt)
	if err != nil{
		//这里面应该写个log  别fmt  你以后上线了怎么办呢
		fmt.Println(err)
		//这里应该不用给用户返回什么信息了  就是报告  服务器内部错误
		return
	}
	tweetId:=*lastRowId
	//todo 这里面 获得数据库的tweetId
	//TODO  这里面去加fanout的逻辑 异步。
	//TODO 先建立一个数据库。
	//TODO 向数据库中去insert数据。
	//TODO if fanout失败 也是whatever了
	userIdInt,err := strconv.Atoi(userId)
	if err != nil{
		fmt.Println(err)
		return
	}
	isCelebrity,err:=DaoRedis.IsCelebrityRedis(userId)
	if err!=nil{
		//	deal error logic
		fmt.Println(err)
		return
	}
	//var intChan =make(chan int)
	var wg sync.WaitGroup

	//----TODO 如果是明星的话 你post twiter之后就啥也别管了  就直接返回
	if isCelebrity==false {

		//var intChan =make(chan int)
		//TODO 想一下异步有多少种方法     可以用go  或者Redis中的任务队列  rabbitMq 但是还是用go吧   服务和服务之间担心有时间差
		fmt.Println("我是不是明星")
		wg.Add(1)
		go utitl.FanOutTweet(int(tweetId), userIdInt)
		wg.Done()
		//这个你无论协程能否成功，都会返回"上传推文成功"
	}
	wg.Wait()





	fmt.Fprintf(w,"上传推文成功")

	//<- intChan
}
