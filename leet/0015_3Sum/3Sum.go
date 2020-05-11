package _015_3Sum

import "sort"

//TODO 1
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := removeDuplicates(nums)
	nums = nums[:length]
	next := tripletGenerator(nums)
	res := make([][]int, 0)
	for {
		a, b, c, ok := next()
		if !ok {
			break
		}
		if a+b+c == 0 {
			res = append(res, []int{a, b, c})
		}
	}
	return res
}

func tripletGenerator(nums []int) func() (a, b, c int, ok bool) {
	idxA, idxB, idxC, ok := 0, 0, 0, true
	if len(nums) < 3 {
		return 0, 0, 0, false
	}
	return nil
}

func removeDuplicates(nums []int) int {
	lastPos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lastPos] {
			nums[lastPos+1] = nums[i]
			lastPos++
		}
	}
	return lastPos + 1
}
