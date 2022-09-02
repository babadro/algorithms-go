package _076_minimum_window_substring

import (
	"math"
)

// tptl. passed. hard. fast solution
func minWindow(s string, t string) string {
	windowStart, matched, subStrStart, minLen := 0, 0, 0, len(s)+1
	freq, exists := [256]int{}, [256]bool{} // single map can be used here, but arrays are faster
	for i := range t {
		char := t[i] - 'A'
		freq[char]++
		exists[char] = true
	}

	for windowEnd := range s {
		rightChar := s[windowEnd] - 'A'
		if exists[rightChar] {
			freq[rightChar]--
			if freq[rightChar] >= 0 {
				matched++
			}
		}

		for matched == len(t) {
			curLen := windowEnd - windowStart + 1
			if minLen > curLen {
				minLen = curLen
				subStrStart = windowStart
			}

			leftChar := s[windowStart] - 'A'
			windowStart++
			if exists[leftChar] {
				if freq[leftChar] == 0 {
					matched--
				}

				freq[leftChar]++
			}
		}
	}

	if minLen > len(s) {
		return ""
	}

	return s[subStrStart : subStrStart+minLen]
}

// passed. hard. 20%, 22% - slow solution
func minWindow2(s string, t string) string {
	pattern, remainingChars := make(map[byte]int, 26), make(map[byte]int, 26)
	for i := range t {
		ch := t[i]
		pattern[ch]++
		remainingChars[ch]++
	}

	freq := make(map[byte]int, 26)
	l := 0
	minLen := math.MaxInt64
	resL, resR := 0, -1
	for r := range s {
		rChar := s[r]
		if pattern[rChar] > 0 {
			freq[rChar]++
			if len(remainingChars) > 0 {
				remainingCount := remainingChars[rChar]
				if remainingCount == 1 {
					delete(remainingChars, rChar)
				} else if remainingCount > 1 {
					remainingChars[rChar]--
				}
			}
		}

		if len(remainingChars) == 0 {
			if curL := r - l + 1; curL < minLen {
				minLen, resL, resR = curL, l, r
			}

			for l < r {
				lChar := s[l]
				patternCount := pattern[lChar]

				if patternCount == 0 || patternCount < freq[lChar] {
					l++
					if patternCount > 0 {
						freq[lChar]--
					}

					if curL := r - l + 1; curL < minLen {
						minLen, resL, resR = curL, l, r
					}

					continue
				}

				break
			}
		}
	}

	return s[resL : resR+1]
}
