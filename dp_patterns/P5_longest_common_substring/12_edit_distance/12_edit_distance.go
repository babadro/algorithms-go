package _2_edit_distance

import "github.com/babadro/algorithms-go/utils"

func findMinOperationsBruteForce(s1, s2 string) int {
	return recBruteForce(s1, s2, 0, 0)
}

func recBruteForce(s1, s2 string, i1, i2 int) int {
	if i1 == len(s1) {
		return len(s2) - i2
	}

	if i2 == len(s2) {
		return len(s1) - i1
	}

	if s1[i1] == s2[i2] {
		return recBruteForce(s1, s2, i1+1, i2+1)
	}

	c1 := 1 + recBruteForce(s1, s2, i1+1, i2)
	c2 := 1 + recBruteForce(s1, s2, i1, i2+1)
	c3 := 1 + recBruteForce(s1, s2, i1+1, i2+1)

	return utils.Min3(c1, c2, c3)
}

func findMinOperationsTopDown(s1, s2 string) int {
	dp := utils.Make2DArr(len(s1), len(s2), -1)

	return recTopDown(dp, s1, s2, 0, 0)
}

func recTopDown(dp [][]int, s1, s2 string, i1, i2 int) int {
	if i1 == len(s1) {
		return len(s2) - i2
	}

	if i2 == len(s2) {
		return len(s1) - i1
	}

	if dp[i1][i2] == -1 {
		if s1[i1] == s2[i2] {
			dp[i1][i2] = recTopDown(dp, s1, s2, i1+1, i2+1)
		} else {
			c1 := 1 + recTopDown(dp, s1, s2, i1+1, i2)
			c2 := 1 + recTopDown(dp, s1, s2, i1, i2+1)
			c3 := 1 + recTopDown(dp, s1, s2, i1+1, i2+1)

			dp[i1][i2] = utils.Min3(c1, c2, c3)
		}
	}

	return dp[i1][i2]
}

func findMinOperationsBottomUp(s1, s2 string) int {
	dp := utils.Make2DArr(len(s1)+1, len(s2)+1, 0)

	for i1 := 0; i1 < len(s1); i1++ {
		dp[i1][0] = i1
	}

	for i2 := 0; i2 < len(s2); i2++ {
		dp[0][i2] = i2
	}

	for i1 := 1; i1 <= len(s1); i1++ {
		for i2 := 1; i2 <= len(s2); i2++ {
			if s1[i1-1] == s2[i2-1] {
				dp[i1][i2] = dp[i1-1][i2-1]
			} else {
				dp[i1][i2] = 1 + utils.Min3(
					dp[i1-1][i2],   // delete
					dp[i1][i2-1],   // insert
					dp[i1-1][i2-1], // replace
				)
			}

		}
	}

	return dp[len(s1)][len(s2)]
}
