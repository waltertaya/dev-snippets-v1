package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) {
	// two pointer solution
	
	reverse := strings.Split(str, "")
	i, j := 0, len(reverse) - 1
	for i <= j {
		temp := reverse[i]
		reverse[i] = reverse[j]
		reverse[j] = temp

		i++
		j--
	}
	fmt.Println(reverse)
}

func main() {
	var str string
	fmt.Print("Enter the string to reverse: ")
	fmt.Scan(&str)

	reverseString(str)
}
