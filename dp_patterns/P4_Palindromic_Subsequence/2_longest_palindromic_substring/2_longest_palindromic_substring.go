package __longest_palindromic_substring

import "github.com/babadro/algorithms-go/utils"

// see leetcode 5
func findLPSLenBruteForce(st string) int {
	return recBruteForce(st, 0, len(st)-1)
}

func recBruteForce(st string, startIDx, endIDx int) int {
	if startIDx > endIDx {
		return 0
	}

	if startIDx == endIDx {
		return 1
	}

	if st[startIDx] == st[endIDx] {
		remainingLen := endIDx - startIDx - 1
		if remainingLen == recBruteForce(st, startIDx+1, endIDx-1) {
			return remainingLen + 2
		}
	}

	c1 := recBruteForce(st, startIDx+1, endIDx)
	c2 := recBruteForce(st, startIDx, endIDx-1)

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

	return recTopDown(st, dp, 0, len(st)-1)
}

func recTopDown(st string, dp [][]int, startIDx, endIDx int) int {
	if startIDx > endIDx {
		return 0
	}

	if startIDx == endIDx {
		return 1
	}

	if dp[startIDx][endIDx] == -1 {
		if st[startIDx] == st[endIDx] {
			remainingLen := endIDx - startIDx - 1
			if remainingLen == recTopDown(st, dp, startIDx+1, endIDx-1) {
				dp[startIDx][endIDx] = remainingLen + 2
				return dp[startIDx][endIDx]
			}
		}

		c1 := recTopDown(st, dp, startIDx+1, endIDx)
		c2 := recTopDown(st, dp, startIDx, endIDx-1)
		dp[startIDx][endIDx] = utils.Max(c1, c2)
	}

	return dp[startIDx][endIDx]
}

func findLPSLenBottomUp(st string) int {
	dp := make([][]bool, len(st))
	for i := range dp {
		dp[i] = make([]bool, len(st))
	}

	for i := 0; i < len(st); i++ {
		dp[i][i] = true // every string with len=1 is a palindrome
	}

	maxLen := 1
	for startIDx := len(st) - 1; startIDx >= 0; startIDx-- {
		for endIDx := startIDx + 1; endIDx < len(st); endIDx++ {
			if st[startIDx] == st[endIDx] {
				// endIDx-startIDx = 1 if len(st[startIDx:endIDx]) == 2, so it is a palindrome
				if endIDx-startIDx == 1 || dp[startIDx+1][endIDx-1] {
					dp[startIDx][endIDx] = true
					maxLen = utils.Max(maxLen, endIDx-startIDx+1)
				}
			}
		}
	}

	return maxLen
}
