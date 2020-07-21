package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	LoginRequest(client)
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	User    *User  `json:"user"`
}

func LoginRequest(client *http.Client) {
	u := &User{
		Email: "ycx@gla.ac.uk",
		Password: "123456",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(u)
	resp, err := client.Post("http://127.0.0.1:60001/", "application/json", buf)
	if err != nil {
		log.Fatalf("Http request failed: #{err}\n")
	}
	defer resp.Body.Close()

	// 格式转换
	var target Response
	decodeErr := json.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		log.Fatalf("无法解析JSON: #{err}\n")
	}
	if target.Code != 200 || target.User == nil {
		log.Fatalf("登陆失败！: #{err}\n")
	}
	fmt.Println(target.User)
}
