// #medium tptl #dp
package _416_partition_equal_subset_sum

// passed #topdown
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	sum /= 2

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}

	return req(nums, 0, sum, dp)
}

func req(nums []int, idx, amount int, dp [][]int) bool {
	if idx == len(nums) || amount < 0 {
		return false
	}

	if amount == 0 {
		return true
	}

	if val := dp[idx][amount]; val != 0 {
		return val == 1
	}

	res := req(nums, idx+1, amount, dp) || req(nums, idx+1, amount-nums[idx], dp)

	if res == true {
		dp[idx][amount] = 1
	} else {
		dp[idx][amount] = 2
	}

	return res
}

// passed #bottomup
func canPartitionBottomUp(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	sum /= 2
	n := len(nums)

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	for s := 1; s <= sum; s++ {
		if nums[0] == s {
			dp[0][s] = true
		}
	}

	for i := 1; i < n; i++ {
		for s := 1; s <= sum; s++ {
			if dp[i-1][s] {
				dp[i][s] = dp[i-1][s]
			} else if s >= nums[i] {
				dp[i][s] = dp[i-1][s-nums[i]]
			}
		}
	}

	return dp[n-1][sum]
}

// best solution. not mine. passed
func canPartitionBottomUpOptimized(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			if j := i - num; j >= 0 && dp[j] {
				dp[i] = true
				if i == sum {
					return true
				}
			}
		}
	}

	return false
}
