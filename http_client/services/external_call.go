package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 10000,
	},
}

func CallLoginOrm(c *gin.Context) {
	LoginRequest(client)
}

func CallLoginSql(c *gin.Context) {
	LoginRequestSql(client)
}

func CallSendMessage(c *gin.Context) {
	SendMessage(client)
}


