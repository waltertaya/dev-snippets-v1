package main

import "fmt"

func main() {
	test1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7} // Expected: 49
	fmt.Println(maxArea(test1))

	test2 := []int{1, 1} // Expected: 1
	fmt.Println(maxArea(test2))
}

func maxArea(height []int) int {
	maxA := 0
	for i, h1 := range height {
		for j := i + 1; j < len(height); j++ {
			h2 := height[j]
			area := min(h1, h2) * (j - i)
			if area > maxA {
				maxA = area
			}
		}
	}
	return maxA
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
