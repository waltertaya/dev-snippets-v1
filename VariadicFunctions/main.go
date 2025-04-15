package main

import "fmt"

func main() {
	sum, avg, count := calculateStats(12, 45, 67, 89, 23)
	fmt.Printf("Sum: %d\nAverage: %.2f\nCount: %d\n", sum, avg, count)

	nums := []int{23, 78, 67, 23, 99, 89, 72, 93}

	sum, avg, count = calculateStats(nums...)
	fmt.Printf("Sum: %d\nAverage: %.2f\nCount: %d\n", sum, avg, count)
}

func calculateStats(nums ...int) (sum int, average float64, count int) {

	for _, value := range nums {
		sum += value
		count++
	}
	average = float64(sum / count)
	return
}
