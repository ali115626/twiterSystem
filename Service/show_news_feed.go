package Service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	Dao "twiter/tweetSystem/Dao/mysql"
)

//Input: userId
//select tweetId from table where  userId= userId order by createAt;
//这个肯定返回很多行
//这里面弄成一个结构体
//[]struct{
//userId:
//content:
//createAt:
//}
//每一行for row.next(){
//select *(userId,content,createAt ) from tweetTable;
//struct=append(struct,userId,content,createAt )
//}
//structMar:=Json.Marshal([]struct)
//Fmt.Printlnf(w,structMar)




func ShowNewsFeed(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	userId := requestMap["userId"][0]

	//todo 这个userId其实也是friendId
	//把string 如何转为int

	userIdInt,err:=strconv.Atoi(userId)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//
	tweetIdList,err:=Dao.QueryTweetIdFromNewsfeedTable(userIdInt)
	fmt.Println(tweetIdList)


	tweetConentList,err:=Dao.FetchDetailFromTweetTable(tweetIdList)

	if err != nil{
		fmt.Println(err)
		return
	}

	tweetConentListMar,err:=json.Marshal(tweetConentList)
	if err!= nil{
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w,string(tweetConentListMar))

}
