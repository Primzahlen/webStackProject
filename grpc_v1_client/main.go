package main

import (
	"google.golang.org/grpc"
	"grpc_v1_client/pbfiles"
	"grpc_v1_client/services"
	"log"
)
func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		log.Fatal("Fail to connect grpc server #{err}\n")
	}

	defer conn.Close()
	userClient := pbfiles.NewUserServiceClient(conn)
	services.OneWayLogin(userClient, b)
	//调用transfer函数
	transferClient := pbfiles.NewTransferServiceClient(conn)
	services.SendMessage(transferClient)
}
