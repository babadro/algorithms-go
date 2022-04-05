// https://leetcode.com/problems/maximum-length-of-repeated-subarray/
package _1_longest_common_substring

import "github.com/babadro/algorithms-go/utils"

func findLCSLen(s1, s2 string) int {
	return recBruteForce(s1, s2, 0, 0, 0)
}

func recBruteForce(s1, s2 string, i1, i2, count int) int {
	if i1 == len(s1) || i2 == len(s2) {
		return count
	}

	if s1[i1] == s2[i2] {
		count = recBruteForce(s1, s2, i1+1, i2+1, count+1)
	}

	c1 := recBruteForce(s1, s2, i1, i2+1, 0)
	c2 := recBruteForce(s1, s2, i1+1, i2, 0)

	return utils.Max3(count, c1, c2)
}

func findLCSLenTopDown(s1, s2 string) int {
	minLen := utils.Min(len(s1), len(s2))
	dp := make([][][]int, len(s1))
	for i := range dp {
		dp[i] = make([][]int, len(s2))
		for j := range dp[i] {
			dp[i][j] = make([]int, minLen)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	return recTopDown(dp, s1, s2, 0, 0, 0)
}

func recTopDown(dp [][][]int, s1, s2 string, i1, i2, count int) int {
	if i1 == len(s1) || i2 == len(s2) {
		return count
	}

	if dp[i1][i2][count] == -1 {
		c1 := count
		if s1[i1] == s2[i2] {
			c1 = recTopDown(dp, s1, s2, i1+1, i2+1, count+1)
		}
		c2 := recTopDown(dp, s1, s2, i1+1, i2, 0)
		c3 := recTopDown(dp, s1, s2, i1, i2+1, 0)
		dp[i1][i2][count] = utils.Max3(c1, c2, c3)
	}

	return dp[i1][i2][count]
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
				maxLen = utils.Max(maxLen, dp[i][j])
			}
		}
	}

	return maxLen
}
