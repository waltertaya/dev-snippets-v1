package main

import "fmt"

func main() {
	test1 := hammingWeight(11)
	test2 := hammingWeight(128)
	test3 := hammingWeight(2147483645)

	fmt.Printf("Test1: %v, Test2: %v, Test3: %v\n", test1, test2, test3)
}

func hammingWeight(n int) int {
    var count int

	for n > 0 {
		if n == 1 {
			count++
			break
		}
		bit := n % 2
		if bit == 1 {
			count++
		}
		n = int(n / 2)
	}

	return count
}
