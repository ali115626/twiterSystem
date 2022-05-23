package main

import (
	"fmt"
	"testing"
	//Dao "twiter/tweetSystem/Dao/mysql"
	"twiter/tweetSystem/Dao/rpc"
)

func TestAdd(t *testing.T) {

	//userId:=1
	//resultStr,err:=rpc.GetFriendList(userId)
	//if err!=nil{
	//
	//}
	//fmt.Println(resultStr)

	userId:=10
	//*message2.FriendsResponseCommonStars,error
	FriendsResponseCommonStars,err:=rpc.GetFriendsCommonStarList(userId)
	fmt.Println(FriendsResponseCommonStars)
	fmt.Println(err)
	//if err!=nil{
	//	resultStr,err:=
	//}
	//fmt.Println(resultStr)
	//etFriendsCommonStarList
}