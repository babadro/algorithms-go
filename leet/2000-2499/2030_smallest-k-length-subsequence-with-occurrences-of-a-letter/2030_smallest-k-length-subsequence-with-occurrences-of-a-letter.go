package _2030_smallest_k_length_subsequence_with_occurrences_of_a_letter

func smallestSubsequenceBruteForce(s string, k int, letter byte, repetition int) string {
	return recBruteForce(s, k, 0, repetition, nil, letter)
}

func recBruteForce(s string, k, idx, repetition int, cur []byte, letter byte) string {
	if idx == len(s) {
		if valid(cur, k, letter, repetition) {
			return string(cur)
		}

		return ""
	}

	res1 := recBruteForce(s, k, idx+1, repetition, append(cur, s[idx]), letter)
	res2 := recBruteForce(s, k, idx+1, repetition, cur, letter)

	return min(res1, res2)
}

func valid(seq []byte, k int, letter byte, repetition int) bool {
	if len(seq) < k {
		return false
	}

	count := 0
	for _, b := range seq {
		if b == letter {
			count++
		}
	}

	return count >= repetition
}

func min(s1, s2 string) string {
	if s1 == "" {
		return s2
	}

	if s2 == "" {
		return s1
	}

	if s1 < s2 {
		return s1
	}

	return s2
}
