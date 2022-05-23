package Service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	Dao "twiter/tweetSystem/Dao/mysql"
)
//只展示自己的朋友圈咋办呢   【4 5  6 8 4 8 9 4 2 3 】
//todo  只需要挑出来自己的内容即可


func ShowTimeLine(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	userId := requestMap["userId"][0]

	//先去查一下这个userId一共有多少文章吧    返回tweetIdlist
	//todo  你发了个啥文章 你自己没点数吗    直接去tweet的文章表去找吧  你把userId 和time建一个联合索引
	//TODO 你这个要返回tweetIdList呀
	//TODO 这个是必须要回表
	userIdInt,err:=strconv.Atoi(userId)
	if err!=nil{
		fmt.Println(err)
		return
	}
	twiterMyselfResult,err:=Dao.SelectTweetContentFromTable(userIdInt)
	if err !=nil{
		fmt.Println(err)
		return
	}
	twiterMyselfResultMar,err:=json.Marshal(twiterMyselfResult)

	if err !=nil{
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w,string(twiterMyselfResultMar))








}
