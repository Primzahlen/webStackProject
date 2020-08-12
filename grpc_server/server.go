package main

import (
	"google.golang.org/grpc"
	"grpc_v1/common"
	"grpc_v1/pbfiles"
	"grpc_v1/services"
	"log"
	"net"
)

func main()  {
	//db := common.InitDB()
	//defer db.Close()
	DB := common.InitMysql()
	defer DB.Close()
	rpcServer := grpc.NewServer()
	pbfiles.RegisterUserServiceServer(rpcServer, new(services.UserService))
	pbfiles.RegisterTransferServiceServer(rpcServer, new(services.Transfer))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Fail to start network listening server #{err}\n")
	}
	panic(rpcServer.Serve(listen))
}
