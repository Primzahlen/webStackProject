package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"http_v1_server/common"
	"log"
	"net/http"
	"net/mail"
)

// User type
type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}

// Response type
type Response struct {
	Message string `json:"message"`
	Code int `json:"code"`
	User *User `json:"user"`
}

// Create Login handler with Gorm
func LoginUserOrm(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u User
	decoder.Decode(&u)
	defer r.Body.Close()

	w.Header().Set("Conent-type", "application/json")
	validationErr := validate(u)
	if validationErr != nil {
		json.NewEncoder(w).Encode(Response{
			Code: 500,
			Message: validationErr.Error(),
		})
		return
	}
	// Gorm操作
	db := common.GetDB()
	email := u.Email
	password := u.Password
	var user User
	db.Where("email = ?", email).Find(&user)
	if user.ID == "" {
		json.NewEncoder(w).Encode(Response{
			Code: 422,
			Message: "用户不存在",
		})
	}
	if user.Password != password {
		json.NewEncoder(w).Encode(Response{
			Code: 400,
			Message: "密码错误",
		})
	}

	// 登陆成功
	fmt.Println("登陆成功，用户为：",user.Email)

	json.NewEncoder(w).Encode(Response{
		Code: 200,
		Message: "OK",
		User: &user,
	})
}

// Create Login handler with SQL
func LoginUserSQL(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u User
	decoder.Decode(&u)
	defer r.Body.Close()

	w.Header().Set("Conent-type", "application/json")
	validationErr := validate(u)
	if validationErr != nil {
		json.NewEncoder(w).Encode(Response{
			Code: 500,
			Message: validationErr.Error(),
		})
		return
	}
	// Gorm操作
	db := common.GetMysql()
	email := u.Email
	password := u.Password
	var user User
	err := db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email,&user.Name ,&user.Password)
	if err != nil {
		log.Fatal("查询发生异常！")
	}
	if user.ID == "" {
		json.NewEncoder(w).Encode(Response{
			Code: 422,
			Message: "用户不存在",
		})
	}
	if user.Password != password {
		json.NewEncoder(w).Encode(Response{
			Code: 400,
			Message: "密码错误",
		})
	}

	// 登陆成功
	fmt.Println("登陆成功，用户为：",user.Email)

	json.NewEncoder(w).Encode(Response{
		Code: 200,
		Message: "OK",
		User: &user,
	})
}

func validate(user User) error {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}
	if len(user.Password) < 6 {
		return errors.New("password is too weak")
	}
	return nil
}
