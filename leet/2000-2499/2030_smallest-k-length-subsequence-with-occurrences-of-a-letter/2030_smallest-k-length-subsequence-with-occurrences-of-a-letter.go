package _2030_smallest_k_length_subsequence_with_occurrences_of_a_letter

func smallestSubsequenceBruteForce(s string, k int, letter byte, repetition int) string {
	return recBruteForce(s, k, 0, repetition, nil, letter)
}

func recBruteForce(s string, k, idx, repetition int, cur []byte, letter byte) string {
	if idx == len(s) {
		if valid(cur, k, letter, repetition) {
			return string(cur)
		}

		return "0"
	}

	res1 := recBruteForce(s, k, idx+1, repetition, append(cur, s[idx]), letter)
	res2 := recBruteForce(s, k, idx+1, repetition, cur, letter)

	return min(res1, res2)
}

func valid(seq []byte, k int, letter byte, repetition int) bool {
	if len(seq) != k {
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
	if s1 == "0" {
		return s2
	}

	if s2 == "0" {
		return s1
	}

	if s1 < s2 {
		return s1
	}

	return s2
}

// todo 1 dp solution doesn't work here. Use monostack:
// https://leetcode.com/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/discuss/1502211/Monostack-sort-of
// solve this problem first as example uf using monostack https://leetcode.com/problems/daily-temperatures/discuss/130532/mono-stack
func smallestSubsequenceTopDown(s string, k int, letter byte, repetition int) string {
	dp := make([]map[string]string, len(s))
	for i := range dp {
		dp[i] = make(map[string]string)
	}

	return recTopDown(s, dp, k, 0, repetition, nil, letter)
}

func recTopDown(s string, dp []map[string]string, k, idx, repetition int, cur []byte, letter byte) string {
	if k-len(cur) > len(s)-idx {
		return "0"
	}

	if idx == len(s) || len(cur) == k {
		if valid(cur, k, letter, repetition) {
			return string(cur)
		}

		return "0"
	}

	if _, ok := dp[idx][string(cur)]; !ok {
		s1 := recTopDown(s, dp, k, idx+1, repetition, append(cur, s[idx]), letter)
		s2 := recTopDown(s, dp, k, idx+1, repetition, cur, letter)

		dp[idx][string(cur)] = min(s1, s2)
	}

	return dp[idx][string(cur)]
}
