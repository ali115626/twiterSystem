package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"twiter/friendshipSystem/Dao/rpc/message2"
)


//先print一下 看这个rpc能不能通   到底能不能通呀
//([]string,error)
func GetFriendsCommonStarList(userId int) (*message2.FriendsResponseCommonStars,error){
	//:9000
	conn, err := grpc.Dial("localhost:9010", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		//return nil,err
	}
	defer conn.Close()

	friendServiceClient := message2.NewFriendShipCommonStarServiceClient(conn)
	request := message2.FetchRequestCommonStar{UserId: int64(userId)}
		//FetchRequest{UserId: int64(userId)}


	response, err :=friendServiceClient.GetFriendsListCommonPersonPlusCelebrity(context.TODO(),&request)
	//response, err := friendServiceClient.GetFriendsList(context.TODO(), &request)
	if err!=nil{
		fmt.Println(err)
		//return nil,err
	}
	//fmt.Println("CommonPersonList=client==",response.CommonPersonList)
	//fmt.Println("StarPersonList =client==",response.StarPersonList)

	return response,nil


}
