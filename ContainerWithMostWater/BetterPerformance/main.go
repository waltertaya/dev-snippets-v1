package main

import "fmt"

func main() {
	test1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7} // Expected: 49
	fmt.Println(maxArea(test1))

	test2 := []int{1, 1} // Expected: 1
	fmt.Println(maxArea(test2))
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxA := 0

	for left < right {
		h := min(height[left], height[right])
		area := h * (right - left)
		if area > maxA {
			maxA = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxA
}
