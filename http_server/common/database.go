package common

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var DB *gorm.DB
// gorm initialize
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
	fmt.Printf("You are connected to the database \n")
	//db.AutoMigrate(&services.User{})
	db.DB().SetMaxIdleConns(100)  // Sets the maximum number of connections in the free connection pool.
	db.DB().SetMaxOpenConns(100)  // Sets the maximum number of open database connections
	db.DB().SetConnMaxLifetime(time.Second * 200)  // Sets the maximum time for a reusable connection
	DB = db
	return db
}

func GetDB() *gorm.DB{
	return DB
}

// Database configuration
const (
	driverName = "mysql"
	host = "127.0.0.1"
	port = "3306"
	database = "grpc_v2"
	username = "root"
	password = "ycx123456"
	charset = "utf8"
)
// Database connection pool
var db *sql.DB

// raw sql lib initialize
func InitMysql() *sql.DB{
	// Build connection: "Username: password @TCP (IP: port)/ database? charset=utf8"
	path := strings.Join([]string{username, ":", password, "@tcp(",host, ":", port, ")/", database, "?charset=",charset}, "")
	db, _ = sql.Open(driverName, path)  // Open the database. The former is the driver name
	db.SetMaxOpenConns(100) // Set the maximum number of database connections
	db.SetMaxIdleConns(100) // Sets the maximum number of idle connections to the database on
	db.SetConnMaxLifetime(time.Second * 200)
	if err := db.Ping(); err != nil{  // Verify the connection
		fmt.Println("opon database fail")
	}
	fmt.Println("connnect success")
	return db
}

func GetMysql() *sql.DB {
	return db
}