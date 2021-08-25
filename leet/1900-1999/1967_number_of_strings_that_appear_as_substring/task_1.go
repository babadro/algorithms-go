package _967_number_of_strings_that_appear_as_substring

import "strings"

// easy
func numOfStrings(patterns []string, word string) int {
	counter := 0
	for i := range patterns {
		if idx := strings.Index(word, patterns[i]); idx != -1 {
			counter++
		}
	}

	return counter
}
