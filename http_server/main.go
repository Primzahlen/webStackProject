package main

import (
	"fmt"
	"http_v1_server/common"
	"http_v1_server/services"
	"log"
	"net/http"

)

func main() {
	// db := common.InitDB()
	// defer db.Close()
	DB := common.InitMysql()
	defer DB.Close()
    // net/http server config
    srv := &http.Server{
		Addr:         ":60001",
		// ReadTimeout:  0.5 * time.Second,
		// WriteTimeout: 0.5 * time.Second,
	}
	// srv.SetKeepAlivesEnabled(false)
	// 监听服务
	http.HandleFunc("/LoginUserOrm",services.LoginUserOrm)
	http.HandleFunc("/LoginUserSQL",services.LoginUserSQL)
	http.HandleFunc("/DataTransmission", services.DataTransmission)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Fail to start network listening server %v\n", err)
	}
	fmt.Println("Http服务开启，监听端口：60001 ....")
}
