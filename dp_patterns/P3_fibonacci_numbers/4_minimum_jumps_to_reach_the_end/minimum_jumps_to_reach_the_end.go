// leetcode 45 https://leetcode.com/problems/jump-game-ii/
package __minimum_jumps_to_reach_the_end

import "math"

func CountMinJumpsBruteForce(jumps []int) int {
	return recBruteForce(jumps, 0)
}

func recBruteForce(jumps []int, idx int) int {
	if idx == len(jumps)-1 {
		return 0
	}

	if jumps[idx] == 0 {
		return math.MaxInt64
	}

	minRes := math.MaxInt64
	for i := 1; i <= jumps[idx] && i+idx < len(jumps); i++ {
		if res := recBruteForce(jumps, idx+i); res != math.MaxInt64 {
			minRes = min(minRes, 1+res)
		}
	}

	return minRes
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func CountMinJumpsTopDown(jumps []int) int {
	dp := make([]int, len(jumps)+1)
	for i := range dp {
		dp[i] = -1
	}

	return recTopDown(dp, jumps, 0)
}

func recTopDown(dp []int, jumps []int, idx int) int {
	if idx == len(jumps)-1 {
		return 0
	}

	if jumps[idx] == 0 {
		return math.MaxInt64
	}

	if dp[idx] == -1 {
		minRes := math.MaxInt64
		for i := 1; i <= jumps[idx] && i+idx < len(jumps); i++ {
			if res := recTopDown(dp, jumps, idx+i); res != math.MaxInt64 {
				minRes = min(minRes, 1+res)
			}
		}

		dp[idx] = minRes
	}

	return dp[idx]
}

func CountMinJumpsBottomUp(jumps []int) int {
	dp := make([]int, len(jumps))
	dp[len(jumps)-1] = 0

	for i := len(jumps) - 2; i >= 0; i-- {
		dp[i] = math.MaxInt64
		for j := 1; j <= jumps[i] && j+i < len(jumps); j++ {
			idx := i + j
			if dp[idx] != math.MaxInt64 {
				dp[i] = min(dp[i], dp[idx]+1)
			}
		}
	}

	return dp[0]
}
