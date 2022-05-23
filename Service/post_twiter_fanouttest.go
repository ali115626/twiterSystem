package Service

import (
	"fmt"
	"net/http"
	"strconv"
	"twiter/tweetSystem/utitl"
)

func PostTweetFanOutTest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	userId := requestMap["userId"][0]
	tweetId:=100

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
	//var intChan =make(chan int)
	//TODO 想一下异步有多少种方法     可以用go  或者Redis中的任务队列  rabbitMq 但是还是用go吧   服务和服务之间担心有时间差
	utitl.FanOutTweetTest(tweetId,userIdInt)
	//这个你无论协程能否成功，都会返回"上传推文成功"
	fmt.Fprintf(w,"上传推文成功")








}
