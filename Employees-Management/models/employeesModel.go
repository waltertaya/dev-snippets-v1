package models

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	FirstName string
	LastName string
	Email string
	ContactNumber string
	Age int
	Dob string
	Salary int
	Address string
}