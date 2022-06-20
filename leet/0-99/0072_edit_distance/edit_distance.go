package _072_edit_distance

import "github.com/babadro/algorithms-go/utils"

// passed. not mine. tptl hard. #dynamic programming. #string
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = 1 + utils.Min3(dp[i][j], dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[n][m]
}

// top down. passed
func minDistance2(word1 string, word2 string) int {
	dp := make([][]int, len(word1))
	for i := range dp {
		dp[i] = make([]int, len(word2))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return rec(dp, word1, word2, 0, 0)
}

func rec(dp [][]int, s1, s2 string, i1, i2 int) int {
	if i1 == len(s1) {
		return len(s2) - i2
	}

	if i2 == len(s2) {
		return len(s1) - i1
	}

	if dp[i1][i2] == -1 {
		if s1[i1] == s2[i2] {
			dp[i1][i2] = rec(dp, s1, s2, i1+1, i2+1)
		} else {
			res1 := rec(dp, s1, s2, i1+1, i2)
			res2 := rec(dp, s1, s2, i1, i2+1)
			res3 := rec(dp, s1, s2, i1+1, i2+1)

			dp[i1][i2] = min3(res1, res2, res3) + 1
		}
	}

	return dp[i1][i2]
}

func min3(a, b, c int) int {
	res := a
	if b < res {
		res = b
	}

	if c < res {
		res = c
	}

	return res
}
