package main

import "fmt"

func fibonacciSeries(n int) {
	// result := [n]int{} => n must be known at compile time(n is part of the type)
	result := make([]int, n)

	prev, curr, temp := 0,1,0

	result[0] = prev
	for i := 1; i < n; i++ {
		result[i] = curr
		temp = curr
		curr += prev
		prev = temp
	}
	fmt.Println(result)
}

func main() {
	var n int

	fmt.Print("Enter the value of n: ")
	fmt.Scan(&n)

	fibonacciSeries(int(n))
}
