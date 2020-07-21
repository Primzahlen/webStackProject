package main

import (
	"fmt"
	"http_v1_server/common"
	"http_v1_server/services"
	"log"
	"net/http"
)

func main() {
	//db := common.InitDB()
	db := common.InitMysql()
	defer db.Close()
	// 监听服务
	//http.HandleFunc("/",services.LoginUserOrm)
	http.HandleFunc("/",services.LoginUserSQL)
	err := http.ListenAndServe(":60001", nil)
	if err != nil {
		log.Fatalf("Fail to start network listening server #{err}\n")
	}

	fmt.Println("Http服务开启，监听端口：60001 ....")
}


