package _154_find_minimum_in_rotated_sorted_array_ii

// tptl. passed. hard
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[l] >= nums[r] {
		// search start left part
		for l < r {
			m := l + (r-l)/2
			if nums[m] > nums[r] {
				l = m + 1
			} else if nums[m] < nums[r] {
				r = m
			} else if nums[r-1] > nums[r] {
				return nums[r]
			} else {
				r--
			}
		}

		return nums[l]
	}

	return nums[0]
}
