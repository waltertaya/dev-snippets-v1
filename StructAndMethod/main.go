package main

import "fmt"

type Person struct {
	name string
	age int
}

func (person *Person) yearsLater(old int) { // pass by reference (*Person)
	person.age += old
}

func (person Person) greet () { // pass by value
	fmt.Printf("Hello %v, you are %v years old\n", person.name, person.age)
}

func main() {
	person := Person {
		name: "Brett",
		age: 21,
	}

	person.greet()

	fmt.Println("Ten years later")

	person.yearsLater(10)
	person.greet()
}
