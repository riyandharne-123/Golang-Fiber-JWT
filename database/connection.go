package database

import (
	"../models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/golang_jwt?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("could not connect to db!")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
