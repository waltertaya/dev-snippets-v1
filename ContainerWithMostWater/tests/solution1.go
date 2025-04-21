package main

import "fmt"

func main() {
	test1 := []int{1,8,6,2,5,4,8,3,7} // ans = 49

	fmt.Println(maxArea(test1))

	test2 := []int{1,1} // ans = 1

	fmt.Println(maxArea(test2))
}

func maxArea(height []int) int {
	maxA := 0
    for i, val := range(height){
		j := len(height) -1
		for j > i {
			area := min(val, height[j]) * ((j - i) + 1)
			if area > maxA {
				maxA = area
			}
			j--
		}
	}
	return maxA
}