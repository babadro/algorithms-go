package _724_find_pivot_index

// passed. easy
func pivotIndex(nums []int) int {
	n, sum := len(nums), 0
	for _, num := range nums {
		sum += num
	}

	left, right := 0, sum
	for i := 0; i < n; i++ {
		right -= nums[i]
		if left == right {
			return i
		}

		left += nums[i]
	}

	return -1
}
