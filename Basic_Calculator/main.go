package main

import "fmt"

func basicCalc(num1 int, num2 int,sign string) int {

	result := 0

	switch sign {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Print("Sign does not exits in the basic calculator\n")
	}

	return result
}

func main() {

	var num1, num2 int
	var sign string

	fmt.Print("Enter the first value: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the operator: ")
	fmt.Scan(&sign)

	fmt.Print("Enter the second the value: ")
	fmt.Scan(&num2)

	result := basicCalc(num1, num2, sign)
	fmt.Println(result)
}
