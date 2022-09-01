package _076_minimum_window_substring

import "math"

// tptl. passed. hard. 20%, 22% - slow solution
// todo 2 find faster solution
func minWindow(s string, t string) string {
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
