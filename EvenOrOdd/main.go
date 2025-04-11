package main

import "fmt"

func evenOrOdd(num int) string {
	var result string
	if num % 2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}

	return result
}

func main() {
	var num int // stores the number var
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	fmt.Printf("%v\n", evenOrOdd(num))
}
