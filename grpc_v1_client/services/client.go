package services

import (
	"context"
	"grpc_v1_client/pbfiles"
	"io"
	"testing"
)

func OneWayLogin(userClient pbfiles.UserServiceClient, b *testing.B) {
	userRes, err := userClient.LoginUser(context.Background(), &pbfiles.User{
		Email: "ycx@gla.ac.uk",
		Name: "Ye Caixu",
		Password: "123456",
	})

	if err != nil{
		b.Fatalf("Failed to request RPC server #{err}\n")
	}

	if userRes == nil || userRes.Code != 200 || userRes.User == nil {
		b.Fatalf("grpc response is wrong: %v", userRes)
	}
	//log.Println(userRes.User)

}

func LoginByStream(userClient pbfiles.UserServiceClient, b *testing.B)  {
	usersReq := make([]*pbfiles.User, 0)
	var email string = "ycx@gla.ac.uk"
	var password string = "123456"
	for i := 0; i <b.N; i++ {
		user := &pbfiles.User{Email: email, Password: password}
		usersReq = append(usersReq, user)
	}
	stream, err := userClient.GetUserStream(context.Background(),
		&pbfiles.UsersRequest{Users: usersReq})
	if err != nil {
		b.Fatalf("请求GRPC服务器失败 %v\n", err)
	}
	for  {
		//usersRes := make([] *pbfiles.UserResponse, 0)
		usersRes, err := stream.Recv()
		// 数据已经全部返回
		if err == io.EOF {
			break
		}
		if err != nil {
			b.Fatalf("读取服务端流失败 err: %v\n", err.Error())
		}
		if usersRes == nil{
			b.Fatalf("grpc response is wrong: %v", usersRes)
		}
	}
}
// using sql server response
func OneWayLoginSql(userClient pbfiles.UserServiceClient, b *testing.B) {
	userRes, err := userClient.LoginUserSql(context.Background(), &pbfiles.User{
		Email: "ycx@gla.ac.uk",
		Name: "Ye Caixu",
		Password: "123456",
	})

	if err != nil{
		b.Fatalf("Failed to request RPC server #{err}\n")
	}

	if userRes == nil || userRes.Code != 200 || userRes.User == nil {
		b.Fatalf("grpc response is wrong: %v", userRes)
	}
	//log.Println(userRes.User)

}

func LoginByStreamSql(userClient pbfiles.UserServiceClient, b *testing.B)  {
	usersReq := make([]*pbfiles.User, 0)
	var email string = "ycx@gla.ac.uk"
	var password string = "123456"
	for i := 0; i <b.N; i++ {
		user := &pbfiles.User{Email: email, Password: password}
		usersReq = append(usersReq, user)
	}
	stream, err := userClient.GetUserStreamSql(context.Background(),
		&pbfiles.UsersRequest{Users: usersReq})
	if err != nil {
		b.Fatalf("请求GRPC服务器失败 %v\n", err)
	}
	for  {
		//usersRes := make([] *pbfiles.UserResponse, 0)
		usersRes, err := stream.Recv()
		// 数据已经全部返回
		if err == io.EOF {
			break
		}
		if err != nil {
			b.Fatalf("读取服务端流失败 err: %v\n", err.Error())
		}
		if usersRes == nil{
			b.Fatalf("grpc response is wrong: %v", usersRes)
		}
	}
}