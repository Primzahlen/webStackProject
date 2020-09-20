package main

import (
	"fmt"
	"http_v1_server/common"
	"http_v1_server/services"
	"log"
	"net/http"

)

func main() {
	// The connection is GORM
	// db := common.InitDB()
	// defer db.Close()

	// The connection is raw sql
	DB := common.InitMysql()
	defer DB.Close()
    // net/http server config
    srv := &http.Server{
		Addr:         ":60001",
		// ReadTimeout:  0.5 * time.Second,
		// WriteTimeout: 0.5 * time.Second,
	}
	// srv.SetKeepAlivesEnabled(false)
	// Listening service
	http.HandleFunc("/LoginUserOrm",services.LoginUserOrm)
	http.HandleFunc("/LoginUserSQL",services.LoginUserSQL)
	http.HandleFunc("/DataTransmission", services.DataTransmission)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Fail to start network listening server %v\n", err)
	}
	fmt.Println("Http service on, listening port:60001 ....")
}
