package main

import (
	"fmt"
	"os"
)

func main() {
	mySlice := os.Args
	for idx, arg := range mySlice {
		fmt.Printf("Argument: %v, Index: %v\n", arg, idx)
	}
}
