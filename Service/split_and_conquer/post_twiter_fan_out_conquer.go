package split_and_conquer

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	Dao "twiter/tweetSystem/Dao/mysql"
	"twiter/tweetSystem/utitl"
)

func PostTweetFanOutSplit(w http.ResponseWriter, r *http.Request) {
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
	var intChan =make(chan int)
	Dao.MysqlConn =Dao.MysqlInit()

	//TODO 想一下异步有多少种方法     可以用go  或者Redis中的任务队列  rabbitMq 但是还是用go吧   服务和服务之间担心有时间差
	go utitl.FanOutTweet(int(tweetId),userIdInt)
	//这个你无论协程能否成功，都会返回"上传推文成功"
	fmt.Fprintf(w,"上传推文成功")

	<- intChan
}




