package _494_target_sum

// doesn't work
func findTargetTopDownWrong(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, total+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDownWrong(dp, nums, 0, 0, target)
}

func recTopDownWrong(dp [][]int, nums []int, idx, cur, target int) int {
	if idx == len(nums) {
		if cur == target {
			return 1
		}

		return 0
	}

	if dp[idx+1][cur] == -1 {
		res1 := recTopDownWrong(dp, nums, idx+1, cur+nums[idx], target)
		res2 := recTopDownWrong(dp, nums, idx+1, cur-nums[idx], target)

		dp[idx+1][cur] = res1 + res2
	}

	return dp[idx+1][cur]
}

func findTargetSumWaysBruteForce(nums []int, target int) int {
	return recBruteForce(nums, 0, 0, target)
}

func recBruteForce(nums []int, idx, cur, target int) int {
	if idx == len(nums) {
		if cur == target {
			return 1
		}

		return 0
	}

	res1 := recBruteForce(nums, idx+1, cur+nums[idx], target)
	res2 := recBruteForce(nums, idx+1, cur-nums[idx], target)

	return res1 + res2
}
