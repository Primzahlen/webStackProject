package services

import (
	"github.com/gin-gonic/gin"
	"net/http"

)
// Provides an interface for external calls, implemented using the NET/HTTP package
var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 5000,
		// DisableKeepAlives: true,
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