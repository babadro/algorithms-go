package __subsets_with_duplicates

import "sort"

// Given a set of numbers that might contain duplicates, find all of its distinct subsets.
// see leetcode 90
func findSubsetsWithDuplicates(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{{}}
	prevEndIDx := 0
	for i := 0; i < len(nums); i++ {
		startIDx := 0
		if i > 0 && nums[i] == nums[i-1] {
			startIDx = prevEndIDx + 1
		}

		prevEndIDx = len(res) - 1
		for j := startIDx; j <= prevEndIDx; j++ {
			newSet := make([]int, len(res[j])+1)
			copy(newSet, res[j])
			newSet[len(newSet)-1] = nums[i]

			res = append(res, newSet)
		}
	}

	return res
}
