package _290_word_pattern

import "strings"

// todo 1 look for solution without len(map)
func wordPattern2(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	m1, m2 := make(map[byte]int), make(map[string]int)
	idx1, idx2, ok := 0, 0, false
	for i := 0; i < len(pattern); i++ {
		if idx1, ok = m1[pattern[i]]; !ok {
			idx1 = len(m1)
			m1[pattern[i]] = idx1
		}

		if idx2, ok = m2[words[i]]; !ok {
			idx2 = len(m2)
			m2[words[i]] = idx2
		}

		if idx1 != idx2 {
			return false
		}
	}

	return true
}

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
