package contest

import "strings"

func numOfStrings(patterns []string, word string) int {
	counter := 0
	for i := range patterns {
		if idx := strings.Index(word, patterns[i]); idx != -1 {
			counter++
		}
	}

	return counter
}
