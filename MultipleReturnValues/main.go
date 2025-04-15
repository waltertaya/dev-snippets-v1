package main

import "fmt"

func main() {
	grades := []float64{56.0, 78.5, 45.0, 90.0, 39.0}

	avg, max, min, passed := analyzeGrades(grades)

	fmt.Printf("Average: %.2f\nHighest: %.2f\nLowest: %.2f\nAll Passed: %v\n",
		avg, max, min, passed)

}

func analyzeGrades(grades []float64) (avg float64, highest float64, lowest float64, allPassed bool) {
	sum := 0.0
	allPassed, highest, lowest = true, grades[0], grades[0]

	for _, value := range grades {
		if value < lowest {
			lowest = value
		}
		if value > highest {
			highest = value
		}
		if value < 50 {
			allPassed = false
		}

		sum += value
	}
	avg = sum / float64(len(grades))

	return
}
