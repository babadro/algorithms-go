package _0581_shortest_unsorted_continuous_subarray

import (
	"math"
	"sort"
)

// tptl passed. medium
func findUnsortedSubarray(nums []int) int {
	l, r := 0, len(nums)-1
	for l < len(nums)-1 && nums[l] <= nums[l+1] {
		l++
	}

	if l == len(nums)-1 {
		return 0
	}

	for r > 0 && nums[r] >= nums[r-1] {
		r--
	}

	minNum, maxNum := math.MaxInt64, math.MinInt64
	for i := l; i <= r; i++ {
		num := nums[i]
		minNum, maxNum = min(minNum, num), max(maxNum, num)
	}

	for l > 0 && nums[l-1] > minNum {
		l--
	}

	for r < len(nums)-1 && nums[r+1] < maxNum {
		r++
	}

	return r - l + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// naive, but works
func findUnsortedSubarray2(nums []int) int {
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(arr)

	i, j := 0, len(nums)-1
	for ; i < len(nums) && nums[i] == arr[i]; i++ {
	}

	for ; j > i && nums[j] == arr[j]; j-- {
	}

	return j - i + 1
}
