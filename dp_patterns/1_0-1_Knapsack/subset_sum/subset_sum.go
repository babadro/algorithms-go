package subset_sum

// tptl #dp
func canPartition(nums []int, sum int) bool {
	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for s := sum; s >= nums[i]; s-- {
			if dp[s-nums[i]] {
				dp[s] = true
			}
		}
	}

	return dp[sum]
}
