package main

import (
	"fmt"
	"strings"
)

func palindromeChecker(str string) bool {
	// two pointer solution
	result := true

	// change the string into array
	arr := strings.Split(str, "")

	left, right := 0, len(arr) - 1

	for left <= right {
		if arr[left] != arr[right] {
			result = false
			break
		}

		left++
		right--
	}

	return result
}

func main() {
	test := "baba"

	fmt.Println(palindromeChecker(test))
}
