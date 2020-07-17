package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"grpc_v1/pbfiles"

)
import _ "github.com/go-sql-driver/mysql"

var DB *gorm.DB

func InitDB() *gorm.DB{
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "grpc_v1"
	username := "root"
	password := "ycx123456"
	charset := "utf8"
	//loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil{
		panic("failed to connect database, err:" + err.Error())
	}

	db.AutoMigrate(&pbfiles.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB{
	return DB
}