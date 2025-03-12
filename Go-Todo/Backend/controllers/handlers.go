package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/waltertaya/Todo/initializers"
	"github.com/waltertaya/Todo/models"
)

// type Task struct {
// 	gorm.Model
// 	UserId int
// 	Title string
// 	Complete bool
// }

func RegisterUser(ctx *gin.Context) {
	var body struct {
		Username string
		Email    string
		Password string
	}

	ctx.Bind(&body)

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User created successfully"})
}

func LoginUser(ctx *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	ctx.Bind(&body)

	user := models.User{}

	result := initializers.DB.Where("email = ? AND password = ?", body.Username, body.Password).First(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User logged in successfully"})
}

func CreateTask(ctx *gin.Context) {
	var body struct {
		Title string
		UserId int
	}

	ctx.Bind(&body)

	task := models.Task{
		Title: body.Title,
		UserId: body.UserId,
	}

	result := initializers.DB.Create(&task)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Error creating task"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Task created successfully"})
}

func UpdateTask(ctx *gin.Context) {
	var body struct {
		Complete bool
	}

	ctx.Bind(&body)

	id := ctx.Param("id")

	task := models.Task{}

	result := initializers.DB.First(&task, id)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Task not found"})
		return
	}

	task.Complete = body.Complete

	initializers.DB.Save(&task)

	ctx.JSON(200, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task := models.Task{}
	result := initializers.DB.First(&task, id)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Task not found"})
		return
	}
	initializers.DB.Delete(&task)
	ctx.JSON(200, gin.H{"message": "Task deleted successfully"})
}

func GetTasks(ctx *gin.Context) {
	id := ctx.Param("id")
	tasks := []models.Task{}
	result := initializers.DB.Where("user_id = ?", id).Find(&tasks)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Tasks not found"})
		return
	}
	ctx.JSON(200, gin.H{"tasks": tasks})
}
