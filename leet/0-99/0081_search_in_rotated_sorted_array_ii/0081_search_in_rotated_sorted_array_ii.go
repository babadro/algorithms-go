package _081_search_in_rotated_sorted_array_ii

// tptl. passed. medium (hard for me)
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	if target == nums[0] || target == nums[r] {
		return true
	}

	if nums[l] >= nums[r] {
		for l < r {
			m := l + (r-l)/2
			if nums[m] > nums[r] {
				l = m + 1
			} else if nums[m] < nums[r] {
				r = m
			} else if nums[r-1] > nums[r] {
				l = r // we've found start of left half. Break loop here
			} else {
				r--
			}
		}

		// l is a start of left part of the array
		if target < nums[0] {
			return binarySearch(nums, l, len(nums)-1, target)
		}

		return binarySearch(nums, 0, l-1, target)
	}

	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, l, r, target int) bool {
	for l <= r {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return true
		}
	}

	return false
}
