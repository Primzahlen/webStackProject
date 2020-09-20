package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

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
	email := GenerateEmail()
	u := &User{
		Email: email,
		Name: "Tony",
		Password: "123456",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(u)
	resp, err := client.Post("http://127.0.0.1:60001/LoginUserOrm", "application/json", buf)
	if err != nil {
		log.Fatalf("Http request failed: %s",err)
	}
	defer resp.Body.Close()
	// Format conversion
	var target Response
	decodeErr := json.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		log.Fatalf("Unable to parse JSON: %v\n", decodeErr)
	}
	if target.Code != 200 || target.User == nil {
		log.Fatalf("login fail！: %v\n",target)
	}
}


func LoginRequestSql(client *http.Client) {
	email := GenerateEmail()
	u := &User{
		Email: email,
		Name: "Tony",
		Password: "123456",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(u)
	resp, err := client.Post("http://127.0.0.1:60001/LoginUserSQL", "application/json", buf)
	if err != nil {
		log.Fatalf("Http request failed: %v\n", err)
	}
	defer resp.Body.Close()
	// Format conversion
	var target Response
	decodeErr := json.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		log.Fatalf("Unable to parse JSON: %v\n", decodeErr)
	}
	if target.Code != 200 || target.User == nil {
		log.Fatalf("login fail ！: %v \n", target)
	}
}

// Randomly generate emails with id [1-1000] "id@stu.gla.ac.uk"
func GenerateEmail() string{
	rand.Seed(time.Now().Unix())
	id := rand.Intn(1000)+1
	fmt.Println(id)
	Email :=  strconv.Itoa(id) + "@stu.gla.ac.uk"
	return Email
}

