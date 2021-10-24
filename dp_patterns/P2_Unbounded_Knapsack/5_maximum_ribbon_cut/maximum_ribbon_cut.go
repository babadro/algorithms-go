package __maximum_ribbon_cut

import "math"

func maximumRibbonCutTopDown(lengths []int, length int) int {
	dp := make([][]int, len(lengths))
	for i := range dp {
		dp[i] = make([]int, length+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, lengths, 0, length)
}

// ok
func recTopDown(dp [][]int, lengths []int, idx, length int) int {
	if length == 0 {
		return 0
	}

	if idx == len(lengths) {
		return math.MinInt64
	}

	if dp[idx][length] == -1 {
		count1 := math.MinInt64
		if lengths[idx] <= length {
			res := recTopDown(dp, lengths, idx, length-lengths[idx])
			if res != math.MinInt64 {
				count1 = res + 1
			}
		}

		count2 := recTopDown(dp, lengths, idx+1, length)

		dp[idx][length] = max(count1, count2)
	}

	return dp[idx][length]
}

func maximumRibbonBottomUp(lengths []int, length int) int {
	dp := make([][]int, len(lengths))
	for i := range dp {
		dp[i] = make([]int, length+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt64
		}
	}

	for i := range dp {
		dp[i][0] = 0
	}

	for i := range dp {
		for l := 1; l <= length; l++ {
			if i > 0 {
				dp[i][l] = dp[i-1][l]
			}

			if l >= lengths[i] {
				res := dp[i][l-lengths[i]]
				if res != math.MinInt64 {
					dp[i][l] = max(dp[i][l], res+1)
				}
			}
		}
	}

	if res := dp[len(lengths)-1][length]; res != math.MinInt64 {
		return res
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
