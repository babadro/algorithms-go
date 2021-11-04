package _045_jump_game_ii

import "math"

// passed. #medium #dp
func jump(jumps []int) int {
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

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// passed
func jumpButtomUp(jumps []int) int {
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

// passed. tptl. best solution
// https://leetcode.com/problems/jump-game-ii/discuss/1526030/O(N)-Greedy-solution-with-explanation
// Greedy solution - we only jump only when we have to
// i.e we cannot reach that index without jumping previously
// When we have to jump, we know we can jump to "furthest"
func jumpBestSolution(nums []int) int {
	var furthest, currPosition, jumps int
	for i := 0; i < len(nums); i++ {
		if i > currPosition {
			jumps++
			currPosition = furthest
		}

		if nums[i]+i > furthest {
			furthest = max(furthest, nums[i]+i)
		}
	}

	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
