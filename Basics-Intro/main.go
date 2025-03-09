package main

import "fmt"

type Cars struct {
	Make string
	Model string
	Engine string
	Price int
}

// Nested struct
// type Person struct {
// 	Name string
// 	Age int
// 	Address Address
// }

// type Address struct {
// 	Street string
// 	City string
// 	State string
// }

type Person struct {
	Name string
	Age int
	Address struct {
		Street string
		City string
		State string
	}
}
// Methods for struct
func (c Cars) carDetails() string {
	return fmt.Sprintf("Car: %v %v, Engine: %v, Price: %v", c.Make, c.Model, c.Engine, c.Price)
}

// Embedding struct
type truck struct {
	Cars
	Price int
}

func main() {
	dataType()
	formaatting_string()
	condition_statements()
	fmt.Println(sub(12, 2, "Subtracting"))

	// firstName, lastName := getNames()
	// fmt.Println("Welcome", firstName, lastName)

	firstName, _ := getNames()
	fmt.Println("Welcome", firstName)

	x, y := getCoords()
	fmt.Println(x, y)

	car1 := Cars{"Toyota", "Corolla", "VVT-i", 2000000}
	fmt.Println(car1)
	fmt.Println(car1.Make)

	// Methods for struct
	fmt.Println(car1.carDetails())

	// Nested struct
	// person1 := Person{"Walter", 36, Address{"123 Main St", "Nairobi", "Nairobi"}}
	// fmt.Println(person1)
	person1 := Person{"Walter", 36, struct{Street string; City string; State string}{"123 Main St", "Nairobi", "Nairobi"}}
	fmt.Println(person1)

	// Anonymous struct
	anon := struct {
		Name string
		Age int
	}{
		Name: "Taya",
		Age: 36,
	}
	fmt.Println(anon)
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

// func sub(x int, y int) int {
// 	return x -y
// }
func sub(x, y int, message string) int {
	fmt.Println(message)
	return x -y
}

func getNames() (string, string) {
	return "Walter", "Taya"
}

func getCoords() (x, y int) {
	// x and y are initialized with zero values
	return // automatically return x and y
}
