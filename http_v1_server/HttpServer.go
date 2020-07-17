package http_v1_server

import (
	"encoding/json"
	"net/http"
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

// Create Login handler
func LoginUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User

}
