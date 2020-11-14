package _290_word_pattern

import "strings"

// todo 1
// bruteforce
func wordPattern(pattern string, s string) bool {
	pattern1 := make([]int, len(pattern))
	m1 := make(map[int32]int)

	var idx int
	var ok bool
	for i, letter := range pattern {

		if idx, ok = m1[letter]; !ok {
			idx = len(m1)
			m1[letter] = idx
		}

		pattern1[i] = idx
	}

	words := strings.Split(s, " ")
	pattern2 := make([]int, 0)
	m2 := make(map[string]int)

	for _, word := range words {

		if idx, ok = m2[word]; !ok {
			idx = len(m2)
			m2[word] = idx
		}

		pattern2 = append(pattern2, idx)
	}

	if len(pattern1) != len(pattern2) {
		return false
	}

	for i := range pattern1 {
		if pattern1[i] != pattern2[i] {
			return false
		}
	}

	return true
}
