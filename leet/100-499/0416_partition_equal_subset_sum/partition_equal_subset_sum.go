package _416_partition_equal_subset_sum

// tle
func canPartitionBruteForce(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	return reqBruteForce(nums, 0, sum/2)
}

func reqBruteForce(nums []int, idx, amount int) bool {
	if idx == len(nums) || amount < 0 {
		return false
	}

	if amount == 0 {
		return true
	}

	return reqBruteForce(nums, idx+1, amount) || reqBruteForce(nums, idx+1, amount-nums[idx])
}
