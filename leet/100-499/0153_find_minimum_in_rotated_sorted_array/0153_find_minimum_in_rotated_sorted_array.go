package _153_find_minimum_in_rotated_sorted_array

// tptl passed
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[l] > nums[r] {
		for l < r {
			m := l + (r-l)/2
			if nums[m] > nums[r] {
				l = m + 1
			} else {
				r = m
			}
		}

		return nums[l]
	}

	return nums[0]
}
