package utitl

import (
	"fmt"
	"time"
	Dao "twiter/tweetSystem/Dao/mysql"
)

func ConcurrencyInsertIntoDatabase(index int,friendTwoArray [][]int,tweetId int,createAt time.Time){


		//这样并发地向数据库中插入数据
		for _, friendId := range friendTwoArray[index] {
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