package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Cannot Connect To Database")
	}
	fmt.Println("Connected Database")

}
