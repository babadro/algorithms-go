package _081_search_in_rotated_sorted_array_ii

// todo 1
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	if nums[l] >= nums[r] {
		for l < r {
			m := l + (r-l)/2
			if nums[m] <= nums[r] {
				r = m
			} else {
				l = m + 1
			}
		}

		startLeft := l
		if startLeft == 0 { // arr of equal numbers like [1,1,1,1,1]
			return target == nums[0]
		}

		if target < nums[0] {
			return binarySearch(nums, startLeft, len(nums)-1, target)
		}

		return binarySearch(nums, 0, startLeft-1, target)
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
