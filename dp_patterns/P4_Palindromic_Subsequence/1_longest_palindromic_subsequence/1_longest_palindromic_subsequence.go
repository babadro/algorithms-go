package __longest_palindromic_subsequence

import "github.com/babadro/algorithms-go/utils"

func findLPSLenBruteForce(st string) int {
	return findLPSLenRec(st, 0, len(st)-1)
}

func findLPSLenRec(st string, startIDx, endIDx int) int {
	if startIDx > endIDx {
		return 0
	}

	if startIDx == endIDx {
		return 1
	}

	if st[startIDx] == st[endIDx] {
		return 2 + findLPSLenRec(st, startIDx+1, endIDx-1)
	}

	c1 := findLPSLenRec(st, startIDx+1, endIDx)
	c2 := findLPSLenRec(st, startIDx, endIDx-1)

	return utils.Max(c1, c2)
}

func findLPSLenTopDown(st string) int {
	dp := make([][]int, len(st))
	for i := range dp {
		dp[i] = make([]int, len(st))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return findLPSLenRecTopDown(dp, st, 0, len(st)-1)
}

func findLPSLenRecTopDown(dp [][]int, st string, startIDx, endIDx int) int {
	if startIDx > endIDx {
		return 0
	}

	if startIDx == endIDx {
		return 1
	}

	if dp[startIDx][endIDx] == -1 {
		if st[startIDx] == st[endIDx] {
			dp[startIDx][endIDx] = 2 + findLPSLenRecTopDown(dp, st, startIDx+1, endIDx-1)
		} else {
			c1 := findLPSLenRecTopDown(dp, st, startIDx+1, endIDx)
			c2 := findLPSLenRecTopDown(dp, st, startIDx, endIDx-1)

			dp[startIDx][endIDx] = utils.Max(c1, c2)
		}
	}

	return dp[startIDx][endIDx]
}
