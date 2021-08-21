package minimum_subset_sum_difference

import "github.com/babadro/algorithms-go/utils"

func minDiff(nums []int) int {
	return rec(nums, 0, 0, 0)
}

func rec(nums []int, idx, sum1, sum2 int) int {
	if idx == len(nums) {
		return utils.Abs(sum1 - sum2)
	}

	diff1 := rec(nums, idx+1, sum1+nums[idx], sum2)
	diff2 := rec(nums, idx+1, sum1, sum2+nums[idx])

	return utils.Min(diff1, diff2)
}

func minDiffTopDown(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, sum+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(nums, dp, 0, 0, 0)
}

func recTopDown(nums []int, dp [][]int, idx, sum1, sum2 int) int {
	if idx == len(nums) {
		return utils.Abs(sum1 - sum2)
	}

	if dp[idx][sum1] == -1 {
		diff1 := recTopDown(nums, dp, idx+1, sum1+nums[idx], sum2)
		diff2 := recTopDown(nums, dp, idx+1, sum1, sum2+nums[idx])

		dp[idx][sum1] = utils.Min(diff1, diff2)
	}

	return dp[idx][sum1]
}

func minDiffBottomUp(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	sum := totalSum / 2

	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	for i := range dp {
		dp[i][0] = true
	}

	if num := nums[0]; num <= sum {
		dp[0][num] = true
	}

	for numIdx := 1; numIdx < len(nums); numIdx++ {
		for curSum := 1; curSum <= sum; curSum++ {
			if dp[numIdx-1][curSum] {
				dp[numIdx][curSum] = true
			} else if curSum >= nums[numIdx] {
				dp[numIdx][curSum] = dp[numIdx-1][curSum-nums[numIdx]]
			}
		}
	}

	var sum1 int
	for s := sum; s >= 0; s-- {
		if dp[len(nums)-1][s] {
			sum1 = s

			break
		}
	}

	sum2 := totalSum - sum1

	return utils.Abs(sum1 - sum2)
}
