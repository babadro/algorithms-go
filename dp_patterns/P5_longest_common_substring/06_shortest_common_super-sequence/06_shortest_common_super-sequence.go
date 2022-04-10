package _6_shortest_common_super_sequence

import "github.com/babadro/algorithms-go/utils"

func findSCSLenBruteForce(s1, s2 string) int {
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
		return 1 + recBruteForce(s1, s2, i1+1, i2+1)
	}

	len1 := 1 + recBruteForce(s1, s2, i1, i2+1)
	len2 := 1 + recBruteForce(s1, s2, i1+1, i2)

	return utils.Min(len1, len2)
}

func findSCSLenTopDown(s1, s2 string) int {
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
	if i1 == len(s1) {
		return len(s2) - i2
	}

	if i2 == len(s2) {
		return len(s1) - i1
	}

	if dp[i1][i2] == -1 {
		if s1[i1] == s2[i2] {
			dp[i1][i2] = 1 + recTopDown(dp, s1, s2, i1+1, i2+1)
		} else {
			len1 := 1 + recTopDown(dp, s1, s2, i1+1, i2)
			len2 := 1 + recTopDown(dp, s1, s2, i1, i2+1)

			dp[i1][i2] = utils.Min(len1, len2)
		}
	}

	return dp[i1][i2]
}

func findSCSLenBottomUp(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for j := range dp {
		dp[j][0] = j
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + utils.Min(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
