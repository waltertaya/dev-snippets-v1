package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Employees-Management/initializers"
)

func init(){
	initializers.LoadEnv()
}

func main() {
	fmt.Println("Starting a server")

	router := gin.Default()

	router.GET("/employees",func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Reached the endpoint",
		})
	})

	router.Run()
}
