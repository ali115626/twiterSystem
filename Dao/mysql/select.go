package Dao

import (
	"database/sql"
	"errors"
	"fmt"
	"twiter/friendshipSystem/common"
	"twiter/tweetSystem/model"
)



//TODO  query的时候能不能让这个数据库保持长连接呢？
//todo 数据库的三次握手和四次挥手的时间耗时
//mSql := "select * from user where useId in "
//rows, _ := db.Query(mSql)
//rows.Close() //这里如果不释放连接到池里，执行5次后其他并发就会阻塞
//fmt.Println("第 ", i)

//defer func() {
//}()
//var toUserId int
// 循环读取数据
//userId,tweetId,content,createAt

func SearchTweetIdListForCommonPersons(commonPersonList []int)([]int,error){
	//var rows *sql.Rows
	var tweetIdList []int
	for _,userId :=range commonPersonList{
		db:=MysqlInit()
		rows, err := db.Query("select tweetId from News_feed where friendId=? order by createAt Desc",userId)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("select tweetId from Tweet error: %v\n", err))
		}
		defer func() {
		rows.Close()
		}()
		//var toUserId int
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
		//fmt.Println("tweetIdList1=======",tweetIdList)
	}
	 // 会释放数据库连接
	 //fmt.Println("tweetIdList2======",tweetIdList)
	return tweetIdList,nil
}


func SearchTweetContentDependOnTweetIdList(tweetIdList []int)(*[]model.TweetDetail,error){
	//var rows *sql.Rows
	var tweetContentList []model.TweetDetail
	//db:=MysqlInit()
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		//return nil,nil,errors.New(fmt.Sprintf("open database error,err=", err))
		//return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}

	for _,tweetId :=range tweetIdList{
		//select userId,tweetId,content,createAt from Tweet where userId = ?", star
		rows, err := db.Query("select userId,tweetId,content,createAt from Tweet where tweetId=?",tweetId)
		if err != nil {
			//return nil, errors.New(fmt.Sprintf("select tweetId from Tweet error: %v\n", err))
		}
		defer func(){
			rows.Close() // 会释放数据库连接
		}()
		var tweetId int
		var userId int
		var content string
		var createAt string
		for rows.Next() {
			err := rows.Scan(&userId,&tweetId,&content,&createAt)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				//return nil, err
			}
			aa:=model.TweetDetail{
				TweetId:  tweetId,
				UserId:   userId,
				Content:  content,
				CreateAt: createAt,
			}
			tweetContentList = append(tweetContentList, aa)
		}
	}

	return &tweetContentList,nil


}


func SearchCommonPersomContentFromTweetTable(commonPersonList []int)(*[]model.TweetDetail,error){
	tweetIdList,err:=SearchTweetIdListForCommonPersons(commonPersonList)

	fmt.Println("tweetIdList=====tweetIdList=======",tweetIdList)
	if err!=nil{
		fmt.Println(err)
	}
	//这里面再返回一个ContentList
	commonPersonTweetList,err:=SearchTweetContentDependOnTweetIdList(tweetIdList)
	if err!=nil{
		fmt.Println(err)
	}
	return commonPersonTweetList,nil
}

//TODO 咱们这里 用一个连接其实就能把所有的数据douquery   db.rows()连接数据库  db.close()关闭数据库   四次挥手其实挺费时间的

//input:startIdList
//	startIdList=[]
//(*[]model.TweetDetail,error)
func SearchStarContentFromTweetTable(startIdList []int)([]model.TweetDetail,error){
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		//return nil,nil,errors.New(fmt.Sprintf("open database error,err=", err))
		//return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}
	//fmt.Println(startIdList)
	// db.Query("select * from Tweet where userId in")
	//这里面你去for 一下startIdList  但是别把row.close()关闭就行
	//你这里一个orderby   排好序的tweet信息就出来的   so easy !  return []model.TweetDetail
	//主要是因为你这个row 还没关呀  没有四次握手的消耗  先这样吧
	var tweetDetailList []model.TweetDetail
	for _,star := range startIdList {
		//fmt.Println("star====",star)
		rows, err := db.Query("select userId,tweetId,content,createAt from Tweet where userId = ?", star)
		if err != nil {
			return nil,errors.New(fmt.Sprintf("select tweetId from Tweet error: %v\n", err))
		}
		var content string
		var createAt string
		var tweetId int
		var userId int
		for rows.Next() {
			err := rows.Scan(&userId,&tweetId, &content, &createAt)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				//return nil,nil ,err
				//return nil, err
				return nil,err
			}
			tweetDetail := model.TweetDetail{
				TweetId:  tweetId,
				UserId:   star,
				Content:  content,
				CreateAt: createAt,
			}
			//哈哈 这个就是我一个朋友的推文
			//common.MapTweetDetail[createAt]=tweetDetail
			tweetDetailList=append(tweetDetailList,tweetDetail)

			//fmt.Println("tweetDetail=====",tweetDetail)
		}
	}
	//fmt.Println("len(tweetDetailList)====",tweetDetailList)
	return tweetDetailList,nil

}

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