package test

import (
	"google.golang.org/grpc"
	"grpc_v1_client/pb"
	"grpc_v1_client/services"
	"testing"
)

func BenchmarkOneWay(b *testing.B) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		b.Fatal("Fail to connect grpc server #{err}\n")
	}

	defer conn.Close()
	userClient := pb.NewUserServiceClient(conn)
	for i:=0;i<b.N;i++{
		services.OneWayLogin(userClient, b)
	}
}

func BenchmarkOneWaySql(b *testing.B) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		b.Fatal("Fail to connect grpc server #{err}\n")
	}

	defer conn.Close()
	userClient := pb.NewUserServiceClient(conn)
	for i:=0;i<b.N;i++{
		services.OneWayLoginSql(userClient, b)
	}
}


func BenchmarkLoginStream(b *testing.B) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		b.Fatal("Fail to connect grpc server #{err}\n")
	}

	defer conn.Close()
	userClient := pb.NewUserServiceClient(conn)
	services.LoginByStream(userClient, b)
}

func BenchmarkLoginStreamSql(b *testing.B) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		b.Fatal("Fail to connect grpc server #{err}\n")
	}

	defer conn.Close()
	userClient := pb.NewUserServiceClient(conn)
	services.LoginByStreamSql(userClient, b)
}