package _521_longest_uncommon_subsequence_I

import "strings"

func findLUSlength(a string, b string) int {
	lenA, lenB := len(a), len(b)
	if a != b {
		if lenA > lenB {
			return lenA
		}
		return lenB
	}
	return -1
}

func findLUSlengthNaive(a string, b string) int {
	if len(a) > len(b) {
		a, b = b, a
	}
	if strings.Contains(a, b) {
		return -1
	}
	return len(b)
}
