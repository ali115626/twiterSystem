package Dao

import (
	"database/sql"
	"errors"
	"fmt"
	"twiter/friendshipSystem/common"
	"twiter/tweetSystem/model"
)


func QueryMyselfTweetDetailFromTweetTable(userId int)([]model.TweetDetail,error){
	//TODO  这个比较friends  比较大众化  你  其实 orderBy 这个是多余的

	//		select * from Tweet where userId=userId
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		//return nil,nil,errors.New(fmt.Sprintf("open database error,err=", err))
		//return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}

	//你这里一个orderby   排好序的tweet信息就出来的   so easy !  return []model.TweetDetail
	rows, err := db.Query("select tweetId,content,createAt from Tweet where userId=? order by createAt Desc", userId)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("select tweetId from Tweet error: %v\n", err))
	}
	var content string
	var createAt string
	var tweetId int
	defer func() {
		rows.Close() // 会释放数据库连接
	}()
	//var createAtList []string
	var tweetDetailList []model.TweetDetail
	for rows.Next() {
		err := rows.Scan(&tweetId, &content, &createAt)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			//return nil,nil ,err
			return nil,err
		}

		//你用不用再去根据userId去拿那个头像呀
		tweetDetail := model.TweetDetail{
			TweetId:  tweetId,
			UserId:   userId,
			Content:  content,
			CreateAt: createAt,
		}
		//哈哈 这个就是我一个朋友的推文
		//common.MapTweetDetail[createAt]=tweetDetail
		tweetDetailList=append(tweetDetailList,tweetDetail)
		//TweetDetailList = append(TweetDetailList, tweetDetail)

	}
	return tweetDetailList,nil



}





//Input: userId
//select tweetId from table where userId= userId order by createAt;
//这个肯定返回很多行

func QueryTweetIdFromNewsfeedTable(friendId int)([]int,error) {

		db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
		if err != nil {
			fmt.Println("open database error,err=", err)
			return nil, errors.New(fmt.Sprintf("open database error,err=", err))

		}
		//你把这段代码测试一下
		rows, err := db.Query("select tweetId from News_feed where friendId=? order by createAt Desc",friendId)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("select tweetId from Tweet error: %v\n", err))
		}

		defer func() {
			rows.Close() // 会释放数据库连接
		}()
		//var toUserId int
	var tweetIdList []int
	var tweetId int
		// 循环读取数据
		for rows.Next() {
			err := rows.Scan(&tweetId)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return nil, err
			}
			tweetIdList = append(tweetIdList, tweetId)
		}
	return tweetIdList,nil
	}




//userId | content                                                                        | createAt


//TODO 然后你根据userId 去找到userName   反正就是表格的连接   反正就是  连接呀连接     这都不好分库分表了






//input tweetIdList

//return []struct

func FetchDetailFromTweetTable(tweetIdList []int)([]model.TweetDetail,error){
	////userId | content
	//| createAt
	var userId int
	var content string
	var createAt string
	var tweetDetail model.TweetDetail
	var tweetDetailList []model.TweetDetail
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}
	for _,tweetId := range tweetIdList{
		fmt.Println("tweetId=",tweetId)
	//	query 数据库吧
		err = db.QueryRow("select userId,content,createAt from Tweet where tweetId=?",tweetId).Scan(&userId,&content,&createAt)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("select tweetId from Tweet error: %v\n", err))
		}

		tweetDetail.TweetId=tweetId
		tweetDetail.UserId=userId
		tweetDetail.CreateAt=createAt
		tweetDetail.Content=content

		tweetDetailList=append(tweetDetailList,tweetDetail)

	}

	return tweetDetailList,nil
}


func SelectTweetContentFromTable(userId int)([]model.TweetDetail,error) {

	//		select * from Tweet where userId=userId
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		//return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}
	rows, err := db.Query("select tweetId,content,createAt from Tweet where userId=? order by createAt Desc", userId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("select tweetId from Tweet error: %v\n", err))
	}
	var content string
	var createAt string
	var tweetId int
	defer func() {
		rows.Close() // 会释放数据库连接
	}()
	//var toUserId int
	var TweetDetailList []model.TweetDetail

	// 循环读取数据
	for rows.Next() {
		err := rows.Scan(&tweetId, &content, &createAt)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil, err
		}
		//你用不用再去根据userId去拿那个头像呀
		tweetDetail := model.TweetDetail{
			TweetId:  tweetId,
			UserId:   userId,
			Content:  content,
			CreateAt: createAt,
		}
		TweetDetailList = append(TweetDetailList, tweetDetail)

	}
	return TweetDetailList,nil

}

////你把这段代码测试一下
//
//	//userId   | int(11)
//	//| content  | text
//	//| createAt | timestamp
//	//你这个应该有很多行吧
//map[creatAt]
//var mapTweetDetail map[string]model.TweetDetail

//担心mapmapTimeList 进不去

func QueryTweetDetailFromTweetTable(userId int)([]string,map[string]model.TweetDetail,error) {

	//TODO  这个比较friends  比较大众化  你  其实 orderBy 这个是多余的

	//		select * from Tweet where userId=userId
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		return nil,nil,errors.New(fmt.Sprintf("open database error,err=", err))
		//return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}
	rows, err := db.Query("select tweetId,content,createAt from Tweet where userId=? order by createAt Desc", userId)
	if err != nil {
		return nil, nil,errors.New(fmt.Sprintf("select tweetId from Tweet error: %v\n", err))
	}
	var content string
	var createAt string
	var tweetId int
	defer func() {
		rows.Close() // 会释放数据库连接
	}()
	var createAtList []string
	for rows.Next() {
		err := rows.Scan(&tweetId, &content, &createAt)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil,nil ,err
		}

		//你用不用再去根据userId去拿那个头像呀
		tweetDetail := model.TweetDetail{
			TweetId:  tweetId,
			UserId:   userId,
			Content:  content,
			CreateAt: createAt,
		}
		//哈哈 这个就是我一个朋友的推文
		common.MapTweetDetail[createAt]=tweetDetail
		createAtList=append(createAtList,createAt)
		//TweetDetailList = append(TweetDetailList, tweetDetail)

	}
	return createAtList,common.MapTweetDetail,nil
}