package main

import (
	"google.golang.org/grpc"
	"grpc_v1/common"
	"grpc_v1/pb"
	"grpc_v1/services"
	"log"
	"net"
)

func main()  {
	// The connection is GORM
	db := common.InitDB()
	defer db.Close()
	// The connection is raw sql
	DB := common.InitMysql()
	defer DB.Close()
	rpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(rpcServer, new(services.UserService))
	pb.RegisterTransferServiceServer(rpcServer, new(services.Transfer))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Fail to start network listening server %v \n",err)
	}
	panic(rpcServer.Serve(listen))
}