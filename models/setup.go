package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "userdata"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var DB *gorm.DB

func InitDb() *gorm.DB {
	DB = connectDB()

	return DB
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : ", err)
		return nil
	}

	DB = database

	DB.AutoMigrate(&User{})

	return DB
}
