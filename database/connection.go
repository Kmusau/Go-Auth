package database

import (
	"fmt"
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DNS = "root:Musau6565.@tcp(127.0.0.1:3306)/goauth?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func ConnectDB() {
	connection, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the Database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
