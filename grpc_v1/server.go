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
	db := common.InitDB()
	defer db.Close()

	rpcServer := grpc.NewServer()
	pbfiles.RegisterUserServiceServer(rpcServer, new(services.UserService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Fail to start network listening server #{err}\n")
	}
	rpcServer.Serve(listen)
}
