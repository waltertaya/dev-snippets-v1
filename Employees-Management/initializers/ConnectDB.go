package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URI")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		fmt.Println("Database connected successfully")
	} else {
		log.Fatal("Error connecting to the database")
	}
}