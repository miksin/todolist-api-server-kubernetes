package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDb() *gorm.DB {
	info := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open("mysql", info)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})

	return db
}

func CloseDb(db *gorm.DB) error {
	return db.Close()
}
