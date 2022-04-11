package _8_longest_repeating_subsequence

import "github.com/babadro/algorithms-go/utils"

func findLRSLenBruteForce(st string) int {
	return recBruteForce(st, 0, 0)
}

func recBruteForce(st string, i1, i2 int) int {
	if i1 == len(st) || i2 == len(st) {
		return 0
	}

	if i1 != i2 && st[i1] == st[i2] {
		return 1 + recBruteForce(st, i1+1, i2+1)
	}

	c1 := recBruteForce(st, i1+1, i2)
	c2 := recBruteForce(st, i1, i2+1)

	return utils.Max(c1, c2)
}

func findLRSLenTopDown(st string) int {
	dp := make([][]int, len(st))
	for i := range dp {
		dp[i] = make([]int, len(st))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, st, 0, 0)
}

func recTopDown(dp [][]int, st string, i1, i2 int) int {
	if i1 == len(st) || i2 == len(st) {
		return 0
	}

	if dp[i1][i2] == -1 {
		if i1 != i2 && st[i1] == st[i2] {
			dp[i1][i2] = 1 + recTopDown(dp, st, i1+1, i2+1)
		} else {
			c1 := recTopDown(dp, st, i1+1, i2)
			c2 := recTopDown(dp, st, i1, i2+1)

			dp[i1][i2] = utils.Max(c1, c2)
		}
	}

	return dp[i1][i2]
}

func findLRSLenBottomUp(st string) int {
	dp := make([][]int, len(st)+1)
	for i := range dp {
		dp[i] = make([]int, len(st)+1)
	}

	for i := 1; i <= len(st); i++ {
		for j := 1; j <= len(st); j++ {
			if st[i-1] == st[j-1] && i != j {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(st)][len(st)]
}
