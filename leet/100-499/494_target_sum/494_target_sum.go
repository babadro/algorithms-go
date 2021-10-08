package _494_target_sum

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

func findTargetSumTopDown(nums []int, target int) int {
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
	if idx == len(nums) {
		if sum == 0 {
			return 1
		}

		return 0
	}

	if sum < 0 {
		return 0
	}

	if dp[idx][sum] == -1 {
		res1 := recTopDown(dp, nums, idx+1, sum-nums[idx])
		res2 := recTopDown(dp, nums, idx+1, sum)

		dp[idx][sum] = res1 + res2
	}

	return dp[idx][sum]
}

// todo 1 doesn't work
func findTargetSumBottomUp(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total < target || target < -total || (total+target)%2 == 1 {
		return 0
	}

	sum := (total + target) / 2
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}

	for i := range dp {
		dp[i][0] = 1
	}

	for s := 0; s <= sum; s++ {
		if nums[0] == s {
			dp[0][s] = 1
		}
	}

	for i := 1; i < n; i++ {
		for s := 0; s <= sum; s++ {
			dp[i][s] = dp[i-1][s]
			if s >= nums[i] {
				dp[i][s] += dp[i-1][s-nums[i]]
			}
		}
	}

	return dp[n-1][sum]
}

// tptl. passed. #medium (hard for me). #dp
// todo 1 need to understand
func findTargetSumBottomUpOptimized(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total < target || target < -total || (total+target)%2 == 1 {
		return 0
	}

	sum := (total + target) / 2

	dp := make([]int, sum+1)
	dp[0] = 1
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	return dp[sum]
}
