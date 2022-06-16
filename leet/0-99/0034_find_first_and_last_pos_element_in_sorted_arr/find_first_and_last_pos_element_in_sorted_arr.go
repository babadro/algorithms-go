package _034_find_first_and_last_pos_element_in_sorted_arr

import "sort"

// tptl. best solution
func searchRange2(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}

	leftRes := l

	l, r = 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return []int{leftRes, r}
}

// using sort.Search
func searchRange(nums []int, target int) []int {
	l := len(nums)
	start := sort.Search(l, func(i int) bool {
		return nums[i] >= target
	})
	if start >= l || nums[start] != target {
		return []int{-1, -1}
	}
	nums = nums[start:]
	end := sort.Search(len(nums), func(i int) bool {
		return nums[i] != target
	})
	end += start - 1
	return []int{start, end}
}
