package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	if DB != nil {
		return
	}
	connStr := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	} else {
		fmt.Println("PostGres Connection has been created!!")
	}
}
