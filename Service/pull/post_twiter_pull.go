package pull

import (
	"fmt"
	"net/http"
	"time"
	Dao "twiter/tweetSystem/Dao/mysql"
)

//这个直接insert到那个表格就可以了
//这个只是没有fanout的那个过程

//直接insert就行了  好友不好友的用以前的那个表格就行了   给以前的表格去排序

func PostTweetPull(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	userId := requestMap["userId"][0]
	content := requestMap["content"][0]
	createAt:=time.Now()
	_,err =Dao.InsertTweet(userId,content,createAt)
	if err != nil{
		//这里面应该写个log  别fmt  你以后上线了怎么办呢
		fmt.Println(err)
		//这里应该不用给用户返回什么信息了  就是报告  服务器内部错误
		return
	}
	fmt.Fprintf(w,"上传推文成功")
}

