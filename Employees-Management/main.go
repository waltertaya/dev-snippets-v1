package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Employees-Management/controllers"
	"github.com/waltertaya/Employees-Management/initializers"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Starting a server")

	router := gin.Default()

	// POST employees
	router.POST("/employees", controllers.PostEmployees)

	router.Run()
}
