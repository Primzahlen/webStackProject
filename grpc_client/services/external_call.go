package services

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"grpc_v1_client/pb"
	"log"
)

var userClient pb.UserServiceClient
var transferClient pb.TransferServiceClient

// Provides an interface for rpc calls, implemented using the grpc framework
func InitGrpc() (*grpc.ClientConn) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Fail to connect grpc server %v \n",err)
	}
	userClient = pb.NewUserServiceClient(conn)
	transferClient = pb.NewTransferServiceClient(conn)
	return conn
}

// Provides an interface for external calls, implemented using the Gin framework
func CallLoginOrm(c *gin.Context) {
	err := OneWayLogin(userClient)
	if err != nil {
		log.Fatalf("RPC server response error %v\n", err)
	}
}

func CallSendMessage(c *gin.Context) {
	SendMessage(transferClient)
}

func CallLoginSql(c *gin.Context) {
	OneWayLoginSql(userClient)
}