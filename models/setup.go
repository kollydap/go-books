//models/setup.go
package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  )
  
var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database !!!")
	}
	err = database.AutoMigrate(&Book{})
	if err != nil{
		return
	}

	DB = database
}
