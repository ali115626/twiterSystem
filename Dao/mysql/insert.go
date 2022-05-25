package Dao

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	Dao "twiter/friendshipSystem/Dao/mysql"

	//"database/sql"

)

func InsertTweet(userId string,content string,createAt time.Time) (*int64,error){



	db,err:=Dao.MysqlInit()
	//db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	//if err != nil {
	//	fmt.Println("open database error,err=", err)
	//	return nil,errors.New(fmt.Sprintf("open database error,err=", err))
	//}
		result, err := db.Exec("insert into Tweet(userId,content,createAt) values(?,?,?)",userId, content,createAt)
	if err != nil {
		fmt.Println("exec failed, err=", err)
		return nil,errors.New(fmt.Sprintf("exec failed, err=", err))
	}
	lastRow,err:=result.LastInsertId()
	if err !=nil{
		return nil,err
	}
	return &lastRow,nil

}


func InsertIntoNewsFeed(friendId int,tweetId int,createAt time.Time)error{
	//你有必要整这么多链接吗
	//一个连接不就行了吗  这里面的这个是长链接吗？
	//db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	//if err != nil {
	//	fmt.Println("open database error,err=", err)
	//	return err
	//}
	db,err:=Dao.MysqlInit()
	_, err = db.Exec("insert into News_feed(friendId,tweetId,createAt) values(?,?,?)",friendId,tweetId,createAt)
	if err != nil{
		fmt.Println("insert failed,",err)
		return err
	}
	db.Close()

	//friendId  | int(11)   | NO   |     | NULL              |                |
	//| tweetId   | int(11)   | YES  |     | NULL              |                |
	//| createAt
	return nil
}



func InsertIntoNewsFeedTest(friendId int,tweetId int,createAt time.Time)error{
	_, err := MysqlConn.Exec("insert into News_feed_test(friendId,tweetId,createAt) values(?,?,?)",friendId,tweetId,createAt)
	if err != nil{
		fmt.Println("insert failed,",err)
		return err
	}

	return nil
}