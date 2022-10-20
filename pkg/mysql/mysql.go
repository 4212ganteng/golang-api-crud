package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Konekdb() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1)/go_awal?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	fmt.Println("Connected to database")
}
