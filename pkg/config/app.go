package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	var err error
	db, err = gorm.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Database connected successfully!")
}

func GetDB() *gorm.DB {
	return db
}
