package _494_target_sum

// todo 1 doesn't work
func findTargetTopDown(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total < target || target < -total || (total+target)%2 == 1 {
		return 0
	}

	sum := (total + target) / 2
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, sum+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, nums, 0, sum)
}

func recTopDown(dp [][]int, nums []int, idx, sum int) int {
	if sum == 0 {
		return 1
	}

	if sum < 0 || idx == len(nums) {
		return 0
	}

	if dp[idx][sum] == -1 {
		res1 := recTopDown(dp, nums, idx+1, sum-nums[idx])
		res2 := recTopDown(dp, nums, idx+1, sum)

		dp[idx][sum] = res1 + res2
	}

	return dp[idx][sum]
}
