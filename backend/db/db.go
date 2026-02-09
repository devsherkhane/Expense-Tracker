package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:9595520628@@tcp(127.0.0.1:3306)/expense_tracker?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("DB Error:", err)
		panic(err)
	}

	DB = db
	fmt.Println("Database connected successfully ✅")
}
