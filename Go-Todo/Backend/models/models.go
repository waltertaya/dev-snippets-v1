package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserId int
	Title string
	Complete bool
}

type User struct {
	gorm.Model
	Username string
	Email string
	Password string
}