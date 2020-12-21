package _1641_count_sorted_vowel_strings

func countVowelStrings2(n int) int {
	return count2(n, 0)
}

// todo 1 understand and rewrite both solutions on this page:
// https://leetcode.com/problems/count-sorted-vowel-strings/discuss/918663/Go-recursion-and-dp
func count2(n, last int) int {
	if n == 0 {
		return 1
	}

	var res int
	for i := last; i < 5; i++ {
		res += count2(n-1, i)
	}

	return res
}

func countVowelStrings3(n int) int {
	// dp[i] means count of count of words ends at a, e, i, o, u
	dp := make([]int, 5)
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < n-1; i++ {
		next := make([]int, 5)

		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				next[j] += dp[k]
			}

		}

		dp = next
	}

	var res int
	for i := range dp {
		res += dp[i]
	}
	return res
}
