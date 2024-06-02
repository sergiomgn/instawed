package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("InstaWed.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database")
		log.Fatal("Failed to connec to Database")
	}
}
