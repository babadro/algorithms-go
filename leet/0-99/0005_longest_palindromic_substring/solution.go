package _005_longest_palindromic_substring

// passed, but slow
func longestPalindrome2(s string) string {
	dp := make([][][3]int, len(s))
	for i := range dp {
		dp[i] = make([][3]int, len(s))
		for j := range dp[i] {
			dp[i][j][0] = -1
		}
	}

	res := recTopDown(dp, s, 0, len(s)-1)
	return s[res[1] : res[2]+1]
}

func recTopDown(dp [][][3]int, s string, start, end int) [3]int {
	if start > end {
		return [3]int{0, start, end}
	}

	if start == end {
		return [3]int{1, start, end}
	}

	if dp[start][end][0] == -1 {
		if s[start] == s[end] {
			remainingLen := end - start - 1
			if res := recTopDown(dp, s, start+1, end-1); res[0] == remainingLen {
				dp[start][end] = [3]int{end - start + 1, start, end}
				return dp[start][end]
			}
		}

		dp[start][end] = recTopDown(dp, s, start, end-1)

		res2 := recTopDown(dp, s, start+1, end)
		if res2[0] > dp[start][end][0] {
			dp[start][end] = res2
		}
	}

	return dp[start][end]
}
