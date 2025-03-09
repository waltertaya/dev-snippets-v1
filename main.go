package main

import "fmt"

func main() {
	// var hello string = "Hello there friend"
	hello := "Hello there friend" // Infered type
	fmt.Println(hello)

	num1 := 2.3
	fmt.Printf("%T", num1)

	const pi = 22/7 // const be known during compile
}
