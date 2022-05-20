package _162_find_peak_element

import "math"

// tptl. passed best solution
func findPeakElementIterative(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func findPeakElementRecursive(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, l, r int) int {
	if l == r {
		return l
	}
	mid := (l + r) / 2
	if nums[mid] > nums[mid+1] {
		return search(nums, l, mid)
	}
	return search(nums, mid+1, r)
}

// O(n) O(1)
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return n - 1
}

// too long
func findPeakElement3(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if g := gradient(m, nums); g == 1 {
			l = m + 1
		} else if g == -1 {
			r = m - 1
		} else {
			return m
		}
	}

	return -1 // impossible
}

func gradient(idx int, nums []int) int {
	l, m, r := val(idx-1, nums), val(idx, nums), val(idx+1, nums)

	if l < m && r < m {
		return 0 // peak
	} else if l < r {
		return 1 // move right
	} else {
		return -1 // move left
	}
}

func val(idx int, nums []int) int {
	if idx == -1 || idx == len(nums) {
		return math.MinInt64
	}

	return nums[idx]
}
