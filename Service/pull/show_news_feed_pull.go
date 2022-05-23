package pull

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"twiter/friendshipSystem/common"
	Dao "twiter/tweetSystem/Dao/mysql"
	"twiter/tweetSystem/Dao/rpc"
	"twiter/tweetSystem/model"
)



func ShowNewsFeedPull(w http.ResponseWriter, r *http.Request) {
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
	friendList,err:=rpc.GetFriendList(userIdStr)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("friendList===ShowNewsFeedPull==",friendList)
	common.MapTweetDetail=make(map[string]model.TweetDetail)

	//createAt:=time.Now()
	var mapTimeList map[string]model.TweetDetail
	var totalCreateAtList []string
	var timeListFromOneFriendId []string
	for _,friendId := range friendList{
		//todo 这里面弄一个数组去insert
		friendIdInt,err:=strconv.Atoi(friendId)
		if err !=nil{
			fmt.Println(err)
		}
		//根据friendId 获取tweetId所对应的全部的推文内容    返回来一个结构体 createAtList=[createAt1,createAt2,createAt3]  map[createAt]=constant.TweetStruct{
		//}
		timeListFromOneFriendId,mapTimeList,err=Dao.QueryTweetDetailFromTweetTable(friendIdInt)
		if err!=nil{
			fmt.Println(err)
			return
		}
		totalCreateAtList=append(totalCreateAtList,timeListFromOneFriendId...)
		//怎么样能把这些时间给merge起来呢
		//	TODO NewsFeedSystem你这个表格还没有开始建立索引
	}
	//var TweetResult []Tweet

	//createAtListSorted=sort(createAtList)
	//for :=range createAtListSorted{
	//	TweetResult=append(TweetResult,)
	//}
	//fmt.Fprintf(w,json.Marshal(TweetResult))
	//totalCreateAtList=totalCreateAtList.([]string)
	 sort.Strings(totalCreateAtList)

	var resultTweetDetailList []model.TweetDetail

	for _,time := range totalCreateAtList{
		//这里面就是整到100就return 吧
		resultTweetDetailList=append(resultTweetDetailList,mapTimeList[time])
		//if i==100{
		//	break
		//}

	}
	resultTweetDetailListMar,err:=json.Marshal(resultTweetDetailList)
	if err!=nil{
		fmt.Println(err)
		return

	}
	fmt.Fprintf(w,string(resultTweetDetailListMar))


}
