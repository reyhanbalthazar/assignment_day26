package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	DB, err = gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}

	DB.AutoMigrate(&Product{}, &Inventory{}, &Order{})
}
