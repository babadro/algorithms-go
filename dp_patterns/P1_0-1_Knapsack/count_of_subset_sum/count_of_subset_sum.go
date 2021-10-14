package count_of_subset_sum

func countSubsetBruteForce(nums []int, sum int) int {
	return recBruteForce(nums, sum, 0)
}

func recBruteForce(nums []int, sum, idx int) int {
	if sum == 0 {
		return 1
	}

	if sum < 0 || len(nums) == 0 || idx == len(nums) {
		return 0
	}

	sum1 := recBruteForce(nums, sum-nums[idx], idx+1)
	sum2 := recBruteForce(nums, sum, idx+1)

	return sum1 + sum2
}

func countSubsetTopDown(nums []int, sum int) int {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, sum+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, nums, sum, 0)
}

func recTopDown(dp [][]int, nums []int, sum, idx int) int {
	if sum == 0 {
		return 1
	}

	if sum < 0 || len(nums) == 0 || idx == len(nums) {
		return 0
	}

	if dp[idx][sum] == -1 {
		sum1 := recTopDown(dp, nums, sum-nums[idx], idx+1)
		sum2 := recTopDown(dp, nums, sum, idx+1)

		dp[idx][sum] = sum1 + sum2
	}

	return dp[idx][sum]
}

func countSubsetBottomUp(nums []int, sum int) int {
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	for s := 1; s <= sum; s++ {
		if nums[0] == s {
			dp[0][s] = 1
		}
	}

	for i := 1; i < n; i++ {
		for s := 1; s <= sum; s++ {
			dp[i][s] = dp[i-1][s]
			if s >= nums[i] {
				dp[i][s] += dp[i-1][s-nums[i]]
			}
		}
	}

	return dp[n-1][sum]
}

func countSubsetBottomUpOptimized(nums []int, sum int) int {
	dp := make([]int, sum+1)
	dp[0] = 1

	for s := 1; s <= sum; s++ {
		if nums[0] == s {
			dp[s] = 1
		}
	}

	for i := 1; i < len(nums); i++ {
		for s := sum; s >= 0; s-- {
			if s >= nums[i] {
				dp[s] += dp[s-nums[i]]
			}
		}
	}

	return dp[sum]
}
