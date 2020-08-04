package services

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"grpc_v1_client/pbfiles"
	"log"
)

var userClient pbfiles.UserServiceClient
var transferClient pbfiles.TransferServiceClient

func InitGrpc() (*grpc.ClientConn) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		log.Fatal("Fail to connect grpc server #{err}\n")
	}
	userClient = pbfiles.NewUserServiceClient(conn)
	transferClient = pbfiles.NewTransferServiceClient(conn)
	return conn
}

func CallLoginOrm(c *gin.Context) {
	OneWayLogin(userClient)
}

func CallSendMessage(c *gin.Context) {
	SendMessage(transferClient)
}

func CallLoginSql(c *gin.Context) {
	OneWayLoginSql(userClient)
}