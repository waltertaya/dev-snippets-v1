package main

import (
	"fmt"
	"strings"
)

func wordFrequencyCounter(str string) map[string]int {
	result := make(map[string]int) // hold result for the answer

	arr := strings.Fields(str)
	for _, val := range arr {
		result[val] = result[val] + 1 // unknown keys return zero
	}

	return result
}

func main() {
	test := "in Go you can count the occurrences of elements in an array using a generic function that works with slices of any type the function count takes a slice and a filter function as arguments and it counts the elements that satisfy the condition defined by the filter function"

	fmt.Println(wordFrequencyCounter(test))
}
