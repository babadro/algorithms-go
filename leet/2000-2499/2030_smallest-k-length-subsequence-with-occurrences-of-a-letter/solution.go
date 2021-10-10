package _2030_smallest_k_length_subsequence_with_occurrences_of_a_letter

// passed tptl #monostack #hard
// https://leetcode.com/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/discuss/1502211/Monostack-sort-of
// todo 1 need to understand
func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	count := 0
	for i := range s {
		if s[i] == letter {
			count++
		}
	}

	extra := count - repetition
	remove := len(s) - k
	var mono, res []byte
	for i := range s {
		top := len(mono) - 1
		for len(mono) > 0 && mono[top] > s[i] && remove > 0 {
			if mono[top] == letter && extra == 0 {
				break
			}

			if mono[top] == letter {
				extra--
			}

			remove--

			mono = mono[:top]
			top--
		}

		mono = append(mono, s[i])
	}

	for i := 0; len(res) < k; i++ {
		if mono[i] != letter && len(res)+max(0, repetition) >= k {
			continue
		}

		res = append(res, mono[i])
		if mono[i] == letter {
			repetition--
		}
	}

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
