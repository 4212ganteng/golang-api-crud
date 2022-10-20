package database

import (
	"fmt"
	"golang-api/models"
	"golang-api/pkg/mysql"
)

func RunMigratation() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("Iam sorry Migration has failed")
	}
	fmt.Println("Migration Success")
}
