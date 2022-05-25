package main

import (
	"fmt"
	"testing"
	Dao "twiter/tweetSystem/Dao/mysql"
)

func TestAdd(t *testing.T) {

	//userId:=1
	//resultStr,err:=rpc.GetFriendList(userId)
	//if err!=nil{
	//
	//}
	//fmt.Println(resultStr)
	//__________________________   ______________________

	//userId:=10
	////*message2.FriendsResponseCommonStars,error
	//FriendsResponseCommonStars,err:=rpc.GetFriendsCommonStarList(userId)
	//fmt.Println(FriendsResponseCommonStars)
	//fmt.Println(err)
	//__________________________   ______________________
	//
	//
	//if err!=nil{
	//	resultStr,err:=
	//}
	//fmt.Println(resultStr)
	//etFriendsCommonStarList

	//startIdList:=[]int{10,4,40}
	//
	//
	//
	//Dao.SeanrchStarContentFromTweetTable(startIdList)
	List:=[]int{4,10}


	tweetIdList,_:=Dao.SearchTweetIdListForCommonPersons(List)

	fmt.Println(tweetIdList)


}