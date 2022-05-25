package main

import (
	"fmt"
	"testing"
	Dao "twiter/tweetSystem/Dao/mysql"
)
//(t *testing.T)

func testFriend(t *testing.T){


	tweetIdList:=[]int{1,2,3,4,5,6,7,8,9,10}

	tweetContent,_:=Dao.SearchTweetContentDependOnTweetIdList(tweetIdList)

	fmt.Println(tweetContent)

}