package utitl

import (
	//"bufio"
	"fmt"
	//"os"
	"strconv"
	"time"
	Dao "twiter/tweetSystem/Dao/mysql"
	"twiter/tweetSystem/Dao/rpc"
)

func FanOutTweet(tweetId int,userId int){
	//TODO 这个错误不用返回  直接打印日志就行

	//去 friend system里面获取 Friends的信息   然后这个就是rpc调用吧  反正whatever
	//todo 明天弄那个grpc框架
	//friendList:=[]int{2,3,4,5,9}
	//userId:=1
	//resultStr,err:=rpc.GetFriendList(userId)
	friendList,err:=rpc.GetFriendList(userId)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//fmt.Println("friendList=====",friendList)
	//mock big Data friendList
	//TODO  构建一个递增的数组friendList     和一个实验用的数据库  你别去insert一些乱七八糟的东西 进数据库    再开一个数据库
	//todo  a := make([]int, max-min+1)
	//    for i := range a {
	//        a[i] = min + i
	//    }
	//这里面有一个时间tag
	timeStrat:=time.Now()

	createAtLocal:=time.Now()

	if err != nil{
		fmt.Println(err)
	}

	for _,friendId:= range friendList{
		//todo 这里面弄一个数组去insert
		friendIdInt,err:= strconv.Atoi(friendId)
		if err != nil{
			fmt.Println(err)
		}
		err =Dao.InsertIntoNewsFeed(friendIdInt,tweetId,createAtLocal)
		if err!=nil{
			fmt.Println("Insert Into News Feed error,friendId=",friendId,err)
		}
		//	TODO NewsFeedSystem你这个表格还没有开始建立索引
		//这里面有一个时间tag 差不多就是第一次和 最后一次时间的差异
		//friendId=friendId+1

	}
	timeElapse:=time.Since(timeStrat)
	fmt.Println("timeElapse========",timeElapse)

	//Dao.MysqlConn.Close()


	/////////////////////////---------- ----------------------------------------


}


//friendIdInt,err:=friendId
//if err !=nil{
//	fmt.Println(err)
//}
//value=value
//如何将int转换为string




