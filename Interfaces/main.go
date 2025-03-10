package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type fullTimeEmployee struct {
	name string
	salary int
}

type contractor struct {
	name string
	hourlyRate int
	hoursWorked int
}

func (ftEmployee fullTimeEmployee) getName() string {
	return ftEmployee.name
}

func (ftEmployee fullTimeEmployee) getSalary() int {
	return ftEmployee.salary
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyRate * c.hoursWorked
}

func printEmployee(e employee) {
	fmt.Printf("Name: %s, Salary: %d\n", e.getName(), e.getSalary())
}

func main() {
	ftEmployee := fullTimeEmployee {
		name: "John Doe",
		salary: 50000,
	}

	c := contractor {
		name: "Jane Doe",
		hourlyRate: 20,
		hoursWorked: 160,
	}

	printEmployee(ftEmployee)
	printEmployee(c)
}