package _459_repeated_substring_pattern

import "strings"

//  best solution. todo 1 need to understand
func repeatedSubstringPattern2(s string) bool {
	return strings.Contains((s + s)[1:2*len(s)-1], s)
}

// passed
func repeatedSubstringPattern(s string) bool {
Loop:
	for n := 1; n <= len(s)/2; n++ {
		if len(s)%n != 0 {
			continue Loop
		}

		subStr := s[:n]
		for j := 0; j <= len(s)-len(subStr); j += len(subStr) {
			if s[j:j+len(subStr)] != subStr {
				continue Loop
			}
		}

		return true
	}

	return false
}
