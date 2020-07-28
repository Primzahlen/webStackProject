package main

import (
	"http_client/services"
	"net/http"
)

func main() {
	client := &http.Client{}
	services.LoginRequest(client)
}
