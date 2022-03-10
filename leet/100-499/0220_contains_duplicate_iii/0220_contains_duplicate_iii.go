package _0220_contains_duplicate_iii

import "sort"

// passed. tptl hard for me
// todo 2 implement singel pass solution like this: https://leetcode.com/problems/contains-duplicate-iii/discuss/432954/Go-4ms-solution-with-comment
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	ids := make([]int, len(nums))
	for i := range ids {
		ids[i] = i
	}

	sort.Slice(ids, func(i, j int) bool {
		return nums[ids[i]] < nums[ids[j]]
	})

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums) && abs(nums[ids[i]]-nums[ids[j]]) <= t; j++ {
			if abs(ids[i]-ids[j]) <= k {
				return true
			}
		}
	}

	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
