package Pull_Plus_Push

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"
	Dao "twiter/tweetSystem/Dao/mysql"
	"twiter/tweetSystem/Dao/rpc"
	"twiter/tweetSystem/model"
)

//show myself timeline 就是在自己Tweet的表中去找吧

//可以把这几个函数都写到一起吧

func ShowTimeLinesPullPlusPush(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	userId := requestMap["userId"][0]

	userIdInt,err:=strconv.Atoi(userId)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//一个是从rpc中获取数据表
	timeStart:=time.Now()



	result,err:=rpc.GetFriendsCommonStarList(userIdInt)
	if err!=nil{
		fmt.Println(err)
	}

	starList:=result.StarPersonList
	commonPersonList :=result.CommonPersonList

	var starIntList []int
	var commonPersonIntList []int

	for _,star:=range starList{
		starInt,err:=strconv.Atoi(star)
		if err!=nil{
			fmt.Println(err)
		}
		starIntList=append(starIntList,starInt)
	}
	//convert string to int
	for _,commonPerson:=range commonPersonList{
		commonPersonInt,err:=strconv.Atoi(commonPerson)
		if err!=nil{
			fmt.Println(err)
		}
		commonPersonIntList=append(commonPersonIntList,commonPersonInt)
	}
	//fmt.Println("starList==client==",starList)
	//fmt.Println("commonPersonList==client==",commonPersonList)
	starTweetContentList,err:=Dao.SearchStarContentFromTweetTable(starIntList)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//fmt.Println("明星的tweet列表为:",starTweetContentList)
	//*[]model.TweetDetail,error
	commonPersonTweetContentList,err:=Dao.SearchCommonPersomContentFromTweetTable(commonPersonIntList)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("普通人tweet的列表为:",commonPersonTweetContentList)

	TweetContentList,err:=SortStarCommonTweetList(starTweetContentList,*commonPersonTweetContentList)
	if err!=nil{
		fmt.Println(err)
	}
	TweetContentListMar,err:=json.Marshal(TweetContentList)

	timeElapse:=time.Since(timeStart)
	fmt.Println("timeElapse=====",timeElapse)
	//timeElapse=rpc获取friendsList +从tweettable 中获取明星的信息 +从news_feed表中获取普通人的信息+对两个表的信息排序

	fmt.Fprintf(w,string(TweetContentListMar))
	//这个弄成一个函数  然后排好序   返回给端上
	//input:starTweetContentList、commonPersonTweetContentList
	//output:ResultSortedContentList
	//////这里面返回了一个tweet 的一个struct1{}
	//SearchStarsTweetFromTweetTable(starList)
	//////这里面返回了一个tweet 的一个struct2{}
	//SearchCommonPersonTweetContentFromTweetTable(commonPersonList)
	//////这里把strcut1和struct2传进去  排个序  返回一个struct1和struct2的排序的结果
	////sortStarsplusCommonPersons()
	//
	//json.Marshal()
	//fmt.Println()
}

func SortStarCommonTweetList(starTweetContentList []model.TweetDetail,commonPersonTweetContentList []model.TweetDetail)([]model.TweetDetail,error){

	var timeList []string
	//var contentListMap map[string]model.TweetDetail

	contentListMap:=make(map[string]model.TweetDetail)
	//
	//fmt.Println("starTweetContentList=====",starTweetContentList)
	//fmt.Println("commonPersonTweetContentList=====",commonPersonTweetContentList)


	for _,content := range starTweetContentList{
		time:=content.CreateAt
		contentListMap[time]=content
		timeList=append(timeList,time)

	}

	for _,content := range commonPersonTweetContentList{
		time:=content.CreateAt
		contentListMap[time]=content
		timeList=append(timeList,time)

	}
	var totalTweetList []model.TweetDetail

	sort.Strings(timeList)
	//先从小到大  在从大到小吧
	for _,time :=range timeList{
		totalTweetList=append(totalTweetList,contentListMap[time])
	}

	return totalTweetList,nil














}


//
//func SearchStarsTweetFromTweetTable(starList []string) (*[]model.TweetDetail,error){
//
//	var starsTweetList []model.TweetDetail
//	//去tweetTable中去找startList里面的内容
//	for _,starId:=range starList{
//		//这里面返回的就是排好序的  Tweet content的struct
//		startTweet :=Dao.SearchStarContentFromTweetTable(starId)
//		starsTweetList=append(starsTweetList,startTweet)
//	}
//
//
//
//
//
//	return nil ,nil




//}

//
//func SearchStarsTweetContent(){
//	result,err:=rpc.GetFriendsCommonStarList()
//	if err!=nil{
//		fmt.Println(err)
//	}
//	starList:=result.StarPersonList
//}
//
//
//func SearchCommonPersonTweetContent(){
//
//}

//
//func search_stars_tweet_content(){
//
//
//
//
//}
