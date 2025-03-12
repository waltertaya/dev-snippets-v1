package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	var err error

	dsn := os.Getenv("DB_URI")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		fmt.Println("Postgresql database created and connected successfully")
	} else {
		log.Fatal("Fail to connect to the database")
	}
}
