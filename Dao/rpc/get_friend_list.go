package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	//"strconv"
	"twiter/friendshipSystem/Dao/rpc/message"
)

func GetFriendList(userId int)([]string,error) {
	//:9000
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	defer conn.Close()
	friendServiceClient := message.NewFriendShipServiceClient(conn)
	request := message.FetchRequest{UserId: int64(userId)}

	response, err := friendServiceClient.GetFriendsList(context.TODO(), &request)
	if err!=nil{
		return nil,err
	}

	//response,err :=orderServiceClient.GetAreaResult(context.TODO(),&request)

	//好友列表为 int64的

	fmt.Println("您的好友列表为：", response)
	return response.FriendList,nil

}



