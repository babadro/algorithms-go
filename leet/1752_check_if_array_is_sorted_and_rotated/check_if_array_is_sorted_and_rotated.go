package _1752_check_if_array_is_sorted_and_rotated

import "sort"

// passed. best solution. tptl
func check2(nums []int) bool {
	found := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if !found {
				found = true
			} else {
				return false
			}
		}
	}

	if !found {
		return true
	}

	return nums[len(nums)-1] <= nums[0]
}

// passed
func check(nums []int) bool {
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			idx = i
			break
		}
	}

	if idx == 0 {
		return true
	}

	arr := make([]int, len(nums))
	n := len(nums[idx:])
	copy(arr, nums[idx:])
	copy(arr[n:], nums[:idx])

	return sort.IntsAreSorted(arr)
}
