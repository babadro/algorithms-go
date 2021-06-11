package sorting

import "math"

func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		idx := findMin(i, nums)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}

func findMin(i int, nums []int) (idx int) {
	min := math.MaxInt64
	for ; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			idx = i
		}
	}

	return idx
}
