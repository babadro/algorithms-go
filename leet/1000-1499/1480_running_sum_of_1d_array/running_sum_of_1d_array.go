package _1480_running_sum_of_1d_array

// passed. easy
func runningSum(nums []int) []int {
	sums := make([]int, len(nums))

	sums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	return sums
}

func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
}
