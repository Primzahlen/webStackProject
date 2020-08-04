package main

import (
	"github.com/gin-gonic/gin"
	"http_client/services"
)

func main()	{
	r := gin.Default()
	r.GET("/login_orm", services.CallLoginOrm)
	r.GET("/login_sql", services.CallLoginSql)
	r.GET("/send_message", services.CallSendMessage)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
