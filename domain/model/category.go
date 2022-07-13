package model

import (
	"category/common"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
)

var DB *gorm.DB

type Category struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	CategoryName string `gorm:"unique;not_null"json:"category_name"`
	CategoryLevel uint32 `json:"category_level"`
	CategoryParent int64 `json:"category_parent"`
	CategoryImage string `json:"category_image"`
	CategoryDescription string `json:"category_description"`

}

func InitMysql(config common.MysqlConfig)  {
	dbUser := config.User
	password := config.Pwd
	host := config.Host
	port := config.Port
	dbName := config.DataBase
	portS := strconv.FormatInt(port, 10)
	connOption := strings.Join([]string{dbUser, ":", password, "@tcp(", host, ":", portS, ")/", dbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	db, err := gorm.Open("mysql", connOption)
	if err != nil {
		fmt.Println("connect err:", err)
		panic(err)
	}
	db.LogMode(true)


	db.SingularTable(true)   //默认不加复数s
	db.DB().SetMaxIdleConns(20)  //设置连接池,空闲
	db.DB().SetMaxOpenConns(100)   // 设置打开最大连接
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migrate()
}

func migrate() {
	DB.Set("gorm.table_options", "ENGINE=InnoDB").AutoMigrate(&Category{})
}