package _724_find_pivot_index

// todo 1
func pivotIndex(nums []int) int {
	n := len(nums)
	sums := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		sums[i] = nums[i] + sums[i+1]
	}

	sum := 0
	for i := 0; i < n; i++ {
		if sum == sums[i+1] {
			return i
		}
		sum += nums[i]
	}

	return -1
}
