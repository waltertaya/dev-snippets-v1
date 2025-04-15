package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filename := os.Args[1]

	f, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
	}

	input := bufio.NewScanner(f)
	// count := make(map[string]int)

	// reading multiple lines and counting their occurrences

	// for input.Scan() {
	// 	line := input.Text()
	// 	// fmt.Println(line)
	// 	count[line]++
	// }
	// fmt.Println(count)

	// Reading just first line and exiting
	input.Scan()
	line := input.Text()
	fmt.Println(line)
}
