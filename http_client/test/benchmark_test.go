package test

import (
	"http_client/services"
	"net/http"
	"testing"
)

var client = &http.Client{}

func BenchmarkLoginRequest(b *testing.B) {
	for i:=0; i<b.N;i++ {
		services.LoginRequest(client)
	}
}

func BenchmarkLoginRequestSql(b *testing.B) {
	for i:=0; i<b.N;i++ {
		services.LoginRequestSql(client)
	}
}

