package main

import "fmt"

func main() {
	result := intToRoman(3749)
	fmt.Println(result) // Expected: MMMDCCXLIX
	fmt.Println(result == "MMMDCCXLIX")

	result = intToRoman(58)
	fmt.Println(result) // Expected: LVIII
	fmt.Println(result == "LVIII")

	result = intToRoman(1994)
	fmt.Println(result) // Expected: MCMXCIV
	fmt.Println(result == "MCMXCIV")
}

func intToRoman(num int) string {
	vals := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}
	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV", "I",
	}

	roman := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			roman += symbols[i]
			num -= vals[i]
		}
	}

	return roman
}
