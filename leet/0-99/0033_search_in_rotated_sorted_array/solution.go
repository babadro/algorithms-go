package _033_search_in_rotated_sorted_array

// tptl. passed. mine. easy to understand. best solution
func search3(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if nums[l] > nums[r] {
		// find start left part of arr
		for l < r {
			m := l + (r-l)/2
			if nums[m] > nums[r] {
				l = m + 1
			} else {
				r = m
			}
		}

		starLeft := l
		if target < nums[0] {
			return binarySearch3(nums, starLeft, len(nums)-1, target)
		}

		return binarySearch3(nums, 0, starLeft-1, target)
	}

	return binarySearch3(nums, 0, len(nums)-1, target)
}

func binarySearch3(nums []int, l, r, target int) int {
	for l <= r {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}
