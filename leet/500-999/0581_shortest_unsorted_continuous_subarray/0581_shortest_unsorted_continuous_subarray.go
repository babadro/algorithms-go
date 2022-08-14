package _0581_shortest_unsorted_continuous_subarray

import (
	"math"
	"sort"

	"github.com/babadro/algorithms-go/utils"
)

// todo 1 doesn't work
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

	minNum, maxNum := math.MaxInt64, math.MaxInt64
	for i := l; i <= r; i++ {
		num := nums[i]
		minNum, maxNum = utils.Min(minNum, num), utils.Max(maxNum, num)
	}

	for l > 0 && nums[l-1] > minNum {
		l--
	}

	for r < len(nums)-1 && nums[r+1] < maxNum {
		r++
	}

	return r - l
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
