package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23")) // Expected : ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations(""))   // Expected : []
	fmt.Println(letterCombinations("2"))  // Expected: ["a", "b", "c"]
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	mapping := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	results := []string{""} // start with an empty combination

	for i := 0; i < len(digits); i++ {
		letters := mapping[digits[i]]
		var temp []string

		for _, prefix := range results {
			for _, letter := range letters {
				temp = append(temp, prefix+string(letter))
			}
		}

		results = temp // update results with the new combinations
	}

	return results
}
