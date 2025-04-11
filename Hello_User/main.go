package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	
	// variable that store the name
	var name string

	// prompt user to enter their name
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	// great the user `Hello, [Name]!`
	fmt.Printf("Hello %v\n", name)

	// second way
	// reader := bufio.NewScanner(os.Stdin)
	// fmt.Print("Enter your DOB: ")
	// reader.Scan()
	// year := reader.Text()
	// age := 2025 - float64(year)
	// fmt.Printf("You are %v years old\n", age)
}
