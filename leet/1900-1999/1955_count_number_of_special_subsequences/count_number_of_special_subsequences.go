package _1955_count_number_of_special_subsequences

// todo 1 tle
func countSpecialSubsequencesTLE(nums []int) int {
	res := 0
	n := len(nums)
	dp := make([]int, n)
	if nums[0] == 0 {
		dp[0] = 1
	}

	for i := 1; i < n; i++ {
		if nums[i] == 0 {
			dp[i] = 1
		}

		for j := 0; j < i; j++ {
			if nums[i] == 0 {
				if nums[j] == 0 {
					dp[i] += dp[j]
				}
			} else if nums[i] == 1 {
				if nums[j] == 0 || nums[j] == 1 {
					dp[i] += dp[j]
				}
			} else if nums[i] == 2 {
				if nums[j] == 1 || nums[j] == 2 {
					dp[i] += dp[j]
				}
			}
		}

		if nums[i] == 2 {
			res += dp[i]
		}
	}

	return res
}
