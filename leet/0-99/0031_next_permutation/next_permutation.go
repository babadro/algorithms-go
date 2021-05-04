package _031_next_permutation

import (
	"math"
	"sort"
)

// doesn't work
func nextPermutationWithSort(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[n-1] {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			sort.Ints(nums[i+1:])

			return
		}
	}

	if sort.SliceIsSorted(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	}) {
		for i = 0; i < n/2; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}

		return
	}

	for i = 1; i < n && nums[i] == nums[0]; i++ {
	}
	if i >= n-1 {
		return
	}

	replaceIdx := i
	min, idx := math.MaxInt32, i
	for ; i < n-1; i++ {
		if nums[i] < min && nums[i] >= nums[0] {
			min, idx = nums[i], i
		}
	}

	nums[replaceIdx], nums[idx] = nums[idx], nums[replaceIdx]

	sort.Ints(nums[replaceIdx+1:])
}
