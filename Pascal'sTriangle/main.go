package main

import "fmt"

func main() {
	test1 := generate(1)
	test2 := generate(5)

	fmt.Printf("Test1: %v\nTest2: %v\n", test1, test2)
}

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		// Each row has i + 1 elements
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1 // First and last elements are always 1

		// Fill the middle values using values from the previous row
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}
