package main

import (
	"github.com/waltertaya/Todo/initializers"
	"github.com/waltertaya/Todo/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Task{})
}