package main

import (
	"github.com/gin-gonic/gin"
	"grpc_v1_client/services"
)

func main() {
	conn := services.InitGrpc()
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.GET("/login_orm", services.CallLoginOrm)
	r.GET("/login_sql", services.CallLoginSql)
	r.GET("/send_message", services.CallSendMessage)
	defer conn.Close()
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}