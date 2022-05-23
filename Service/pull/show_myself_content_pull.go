package pull

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	Dao "twiter/tweetSystem/Dao/mysql"
)

//func show_myself_content_pull(){

func ShowMyselfContentPull(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	userId := requestMap["userId"][0]
	userIdStr,err:=strconv.Atoi(userId)
	if err!=nil{
		fmt.Println(err)
		return
	}

	//TODO 你这个整一个人的就行了
	//friendIdInt,err:=strconv.Atoi(friendId)
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//根据friendId 获取tweetId所对应的全部的推文内容    返回来一个结构体 createAtList=[createAt1,createAt2,createAt3]  map[createAt]=constant.TweetStruct{
	//}
	//TODO 这个就是一个人发的tweet


	tweetContentList,err:=Dao.QueryMyselfTweetDetailFromTweetTable(userIdStr)
	if err!=nil{
		fmt.Println(err)
		return
	}
	tweetContentListMar,err:=json.Marshal(tweetContentList)

	fmt.Fprint(w,string(tweetContentListMar))

}
