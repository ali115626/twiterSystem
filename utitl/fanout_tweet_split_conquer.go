package utitl

import (
	"bufio"
	"fmt"
	"os"
	"time"
	Dao "twiter/tweetSystem/Dao/mysql"
)

//
func FanOutTweet(tweetId int,userId int){
	//TODO 这个错误不用返回  直接打印日志就行

	//去 friend system里面获取 Friends的信息   然后这个就是rpc调用吧  反正whatever
	//todo 明天弄那个grpc框架

	//TODO  构建一个递增的数组friendList     和一个实验用的数据库  你别去insert一些乱七八糟的东西 进数据库    再开一个数据库
	//todo  a := make([]int, max-min+1)
	//    for i := range a {
	//        a[i] = min + i
	//    }
	amount:=1000
	friendList:=make([]int,amount)

	createAt:=time.Now()
	//这里面有一个时间tag
	timeStrat:=time.Now()
	n:=len(friendList)
	if n >100 {
	//	分而治之
	//	整200个协程咋样
	//	interval:=n/20
		interval:=n/40
		var friendTwoArray [][]int
		i := 1
		for i <= interval {
			//aa:=friendList[i-1:i]
			friendTwoArray=append(friendTwoArray,friendList[(i-1)*10:i*10])
			i=i+1
		}
		//得到了一个二维数组
		friendTwoArray=append(friendTwoArray,friendList[(i-1)*10:])
		for index:=range friendTwoArray{
			go ConcurrencyInsertIntoDatabase(index,friendTwoArray,tweetId,createAt)
		}



	} else {
		//如果小于100万的话 就直接遍历  不用起go协程了
		for _, friendId := range friendList {
			//todo 这里面弄一个数组去insert
			//friendIdInt,err:=friendId
			//if err !=nil{
			//	fmt.Println(err)
			//}
			err := Dao.InsertIntoNewsFeed(friendId, tweetId, createAt)
			if err != nil {
				fmt.Println("Insert Into News Feed error,friendId=", friendId, err)
			}
			//	TODO NewsFeedSystem你这个表格还没有开始建立索引
			//这里面有一个时间tag 差不多就是第一次和 最后一次时间的差异

		}

	}

	timeElapse:=time.Since(timeStrat)
/////////////////////////---------- ----------------------------------------
	fmt.Println("timeElapse========",timeElapse)

	filePath := "./FanOut_try_time.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	amountStr:=fmt.Sprintf("amount=%v\n",amount)
	write.WriteString(amountStr)
	//Flush将缓存的文件真正写入到文件中
	timeElapseStr:=fmt.Sprintf("timeElapse=%v\n",timeElapse)

	write.WriteString(timeElapseStr)

	write.Flush()



	//for friend := range friendList{

	//userID=friend

	//Dao层
	//  insert userId and tweetId into NewsFeedSystem
	//先将就这样理解吧
}
