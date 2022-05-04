package _3_longest_substring_with_maximum_k_distinct_characters

import "github.com/babadro/algorithms-go/utils"

// Given a string, find the length of the longest substring in it with no more than K distinct characters.
// leetcode https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/ locked
func findLength(str string, k int) int {
	maxLen, m, start := 0, make(map[byte]int), 0
	for i := 0; i < len(str); i++ {
		m[str[i]]++
		for len(m) > k {
			startChar := str[start]
			m[startChar]--
			if m[startChar] == 0 {
				delete(m, startChar)
			}

			start++
		}

		maxLen = utils.Max(maxLen, i-start+1)
	}

	return maxLen
}
