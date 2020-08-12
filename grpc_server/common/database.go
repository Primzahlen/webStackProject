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
	fmt.Println("connnect success")
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 100)
	//db.AutoMigrate(&pbfiles.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB{
	return DB
}

//数据库配置
const (
	driverName = "mysql"
	host = "127.0.0.1"
	port = "3306"
	database = "grpc_v1"
	username = "root"
	password = "ycx123456"
	charset = "utf8"
)
//Db数据库连接池
var db *sql.DB

//注意方法名大写，就是public
func InitMysql() *sql.DB{
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{username, ":", password, "@tcp(",host, ":", port, ")/", database, "?charset=",charset}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, _ = sql.Open(driverName, path)
	//设置数据库最大连接数
	db.SetMaxOpenConns(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(time.Second * 100)
	//验证连接
	if err := db.Ping(); err != nil{
		fmt.Println("opon database fail")
	}
	fmt.Println("connect success")
	return db
}

func GetMysql() *sql.DB {
	return db
}