package main

import "fmt"

func main() {
	dataType()
	formaatting_string()
	condition_statements()
}

func dataType() {
	// var hello string = "Hello there friend"
	hello := "Hello there friend" // Infered type
	fmt.Println(hello)

	num1 := 2.3
	fmt.Printf("%T", num1)

	const pi = 22/7 // const be known during compile
}

func formaatting_string() {
	// provide Printf & Sprintf
	fmt.Printf("\nHello, %v", "Taya\n")
	// Formats
	// %v -> used for all
	// %d -> based 10 integer printing
	// %s -> string
	// %T -> print type
	// %f -> float64
}

func condition_statements() {
	// age := 36
	// if age < 18 {
	// 	fmt.Println("You are a child")
	// } else if age > 17 && age < 36 {
	// 	fmt.Println("You are youth adult")
	// } else {
	// 	fmt.Println("You are adult")
	// }

	// Shorthand of the above
	if age := 36; age < 18 {
		fmt.Println("You are a child")
	} else if age > 17 && age < 36 {
		fmt.Println("You are youth adult")
	} else {
		fmt.Println("You are adult")
	}
}