package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Todo/controllers"
	"github.com/waltertaya/Todo/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}


func main() {
	router := gin.Default()

	fmt.Printf("Starting the server at: http://localhost:%s", os.Getenv("PORT"))

	// Create user
	router.POST("/api/v1/register", controllers.RegisterUser)

	// Login user
	router.POST("/api/v1/login", controllers.LoginUser)

	// Create task
	router.POST("/api/v1/task", controllers.CreateTask)

	// Update status of task
	router.PUT("/api/v1/task/:id", controllers.UpdateTask)

	// Delete task
	router.DELETE("/api/v1/task/:id", controllers.DeleteTask)

	// Get all tasks for a user by user id
	router.GET("/api/v1/tasks/user/:id", controllers.GetTasks)

	router.Run()
}
