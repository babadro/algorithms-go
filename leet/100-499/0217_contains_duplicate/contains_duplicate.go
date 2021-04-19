package _217_contains_duplicate

import "sort"

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = true
	}
	return false
}

func containsDuplicateNaive(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
