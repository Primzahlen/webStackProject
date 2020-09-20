package services

import (
	"context"
	"fmt"
	"grpc_v1_client/pb"
	"io"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// Create Login stub with Gorm
func OneWayLogin(userClient pb.UserServiceClient) error{
	email := GenerateEmail()
	userRes, err := userClient.LoginUser(context.Background(), &pb.User{
		Email: email,
		Name: "Tony",
		Password: "123456",
	})
	if err != nil{
		log.Fatalf("Failed to request RPC server %v\n", err)
		return err
	}
	if userRes == nil || userRes.Code != 200 || userRes.User == nil {
		log.Fatalf("grpc response is wrong: %v", userRes)
		return err
	}
	return nil
}

func LoginByStream(userClient pb.UserServiceClient, b *testing.B)  {
	usersReq := make([]*pb.User, 0)
	var email string = "ycx@gla.ac.uk"
	var password string = "123456"
	for i := 0; i <b.N; i++ {
		user := &pb.User{Email: email, Password: password}
		usersReq = append(usersReq, user)
	}
	stream, err := userClient.GetUserStream(context.Background(),
		&pb.UsersRequest{Users: usersReq})
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
			log.Fatalf("读取服务端流失败 err: %v\n", err.Error())
		}
		if usersRes == nil{
			log.Fatalf("grpc response is wrong: %v", usersRes)
		}
	}
}
// // Create Login stub with sql lib
func OneWayLoginSql(userClient pb.UserServiceClient) {
	email := GenerateEmail()
	userRes, err := userClient.LoginUserSql(context.Background(), &pb.User{
		Email: email,
		Name: "Tony",
		Password: "123456",
	})
	if err != nil{
		log.Fatalf("Failed to request RPC server %v\n", err)
	}
	if userRes == nil || userRes.Code != 200 || userRes.User == nil {
		log.Fatalf("grpc response is wrong: %v", userRes)
	}
}

func LoginByStreamSql(userClient pb.UserServiceClient, b *testing.B)  {
	usersReq := make([]*pb.User, 0)
	var email string = "ycx@gla.ac.uk"
	var password string = "123456"
	for i := 0; i <b.N; i++ {
		user := &pb.User{Email: email, Password: password}
		usersReq = append(usersReq, user)
	}
	stream, err := userClient.GetUserStreamSql(context.Background(),
		&pb.UsersRequest{Users: usersReq})
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

// Randomly generate emails with id [1-1000] "id@stu.gla.ac.uk"
func GenerateEmail() string{
	rand.Seed(time.Now().Unix())
	id := rand.Intn(1000)+1
	fmt.Println(id)
	Email :=  strconv.Itoa(id) + "@stu.gla.ac.uk"
	return Email
}