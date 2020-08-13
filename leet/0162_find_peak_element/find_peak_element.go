package _162_find_peak_element

import "math"

// TODO 1 rewrite clean bruteforce from solution and then look O(log2N) solution
// Bruteforce
func findPeakElement2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	left, middle, right := math.MinInt64, nums[0], nums[1]
	i := 0
	for ; i < n; i++ {
		if middle > left && middle > right {
			break
		}
		left = nums[i]
		if middleIdx := i + 1; middleIdx < n {
			middle = nums[middleIdx]
		}
		if rightIdx := i + 2; rightIdx < n {
			right = nums[rightIdx]
		} else {
			right = math.MinInt64
		}
	}
	return i
}
