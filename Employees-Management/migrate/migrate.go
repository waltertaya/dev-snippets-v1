package main

import (
	"github.com/waltertaya/Employees-Management/initializers"
	"github.com/waltertaya/Employees-Management/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Employees{})
}
