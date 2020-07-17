package services

import (
	"context"
	"fmt"
	"grpc_v1_client/pbfiles"
	"io"
	"log"
)

func OneWayLogin(userClient pbfiles.UserServiceClient) {
	userRes, err := userClient.LoginUser(context.Background(), &pbfiles.User{
		Email: "ycx@gla.ac.uk",
		Name: "Ye Caixu",
		Password: "123456",
	})

	if err != nil{
		log.Fatalf("Failed to request RPC server #{err}\n")
	}

	if userRes == nil || userRes.Code != 200 || userRes.User == nil {
		log.Fatalf("grpc response is wrong: %v", userRes)
	}
	fmt.Println(userRes.User)
}

func LoginByStream(userClient pbfiles.UserServiceClient)  {
	usersReq := make([]*pbfiles.User, 0)
	var email string = "ycx@gla.ac.uk"
	var password string = "123456"
	var i int32
	for i = 0; i < 6; i++ {
		user := &pbfiles.User{Email: email, Password: password}
		usersReq = append(usersReq, user)
	}
	stream, err := userClient.GetUserStream(context.Background(),
		&pbfiles.UsersRequest{Users: usersReq})
	if err != nil {
		log.Fatalf("请求GRPC服务器失败 %v\n", err)
	}
	for  {
		//usersRes := make([] *pbfiles.UserResponse, 0)
		usersRes, err := stream.Recv()
		// 数据已经全部返回
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("读取服务端流失败 err: %v\n", err.Error())
		}
		fmt.Println(usersRes)

		//for _, u := range usersRes.UserRep{
		//	if u.User != nil{
		//		fmt.Println(u.User)
		//	}
		//}
	}
}