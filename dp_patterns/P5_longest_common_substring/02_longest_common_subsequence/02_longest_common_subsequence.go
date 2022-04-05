package _2_longest_common_subsequence

import "github.com/babadro/algorithms-go/utils"

func findLCSLen(s1, s2 string) int {
	return recBruteForce(s1, s2, 0, 0)
}

func recBruteForce(s1, s2 string, i1, i2 int) int {
	if i1 == len(s1) || i2 == len(s2) {
		return 0
	}

	if s1[i1] == s2[i2] {
		return 1 + recBruteForce(s1, s2, i1+1, i2+1)
	}

	c1 := recBruteForce(s1, s2, i1+1, i2)
	c2 := recBruteForce(s1, s2, i1, i2+1)

	return utils.Max(c1, c2)
}

func findLCSLenTopDown(s1, s2 string) int {
	dp := make([][]int, len(s1))
	for i := range dp {
		dp[i] = make([]int, len(s2))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, s1, s2, 0, 0)
}

func recTopDown(dp [][]int, s1, s2 string, i1, i2 int) int {
	if i1 == len(s1) || i2 == len(s2) {
		return 0
	}

	if dp[i1][i2] == -1 {
		if s1[i1] == s2[i2] {
			dp[i1][i2] = 1 + recTopDown(dp, s1, s2, i1+1, i2+1)
		} else {
			c1 := recTopDown(dp, s1, s2, i1+1, i2)
			c2 := recTopDown(dp, s1, s2, i1, i2+1)

			dp[i1][i2] = utils.Max(c1, c2)
		}
	}

	return dp[i1][i2]
}

func findLCSLenBottomUp(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	maxLen := 0
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i][j-1])
			}

			maxLen = utils.Max(maxLen, dp[i][j])
		}
	}

	return maxLen
}
