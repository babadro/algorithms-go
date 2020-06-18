package _078_subsets

import "sort"

// TODO 1
func subsets(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{{}}
	}
	if n == 1 {
		return [][]int{{}, nums}
	}
	sort.Ints(nums)
	totalCount := n * n
	combinations := make([][]int, 0, totalCount)
	combinations = append(combinations, []int{})
	start, end := 0, 0
	var next func() (int, bool)
	for i := start; i <= end; i++ {
		next = exceptIterator(combinations[i], nums)
		counter := 0
		for {
			num, ok := next()
			if !ok {
				break
			}
			length := len(combinations[i]) + 1
			combination := make([]int, length)
			copy(combination, combinations[i])
			combination[length-1] = num
			combinations = append(combinations, combination)
			counter++
		}
		if len(combinations) == totalCount {
			break
		}
		start, end = end+1, end+counter
	}
	return combinations
}

func exceptIterator(combination, nums []int) func() (int, bool) {
	combIdx, numsIdx := 0, 0
	return func() (int, bool) {
		for numsIdx < len(nums) {
			if combIdx < len(combination) {
				if nums[numsIdx] < combination[combIdx] {
					numsIdx++
					continue
				}
				if nums[numsIdx] == combination[combIdx] {
					numsIdx++
					combIdx++
					continue
				}
			}
			res := nums[numsIdx]
			numsIdx++
			return res, true
		}
		return 0, false
	}
}
