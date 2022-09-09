package _198_house_robber

// todo2 bottom up

func robBruteForce(nums []int) int {
	return recBruteForce(nums, 0)
}

func recBruteForce(nums []int, i int) int {
	if i >= len(nums) {
		return 0
	}

	res1 := nums[i] + recBruteForce(nums, i+2)
	res2 := recBruteForce(nums, i+1)

	return max(res1, res2)
}

// tptl. passed
func robTopDown(nums []int) int {
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = -1
	}

	return recTopDown(dp, nums, 0)
}

func recTopDown(dp, nums []int, i int) int {
	if i >= len(nums) {
		return 0
	}

	if dp[i] == -1 {
		res1 := nums[i] + recTopDown(dp, nums, i+2)
		res2 := recTopDown(dp, nums, i+1)

		dp[i] = max(res1, res2)
	}

	return dp[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
