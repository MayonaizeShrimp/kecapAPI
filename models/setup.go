package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//connection string
	dsn := "root:@tcp(127.0.0.1:3306)/kecapdb"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//check database connections
	if err != nil {
		panic("failed to connect to database!")
	}

	//migrate models
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Kecap{})

	//assign to higher scope so it can be accessed by models
	DB = database
}
