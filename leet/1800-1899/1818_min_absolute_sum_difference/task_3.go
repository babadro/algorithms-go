package _1818_min_absolute_sum_difference

import (
	"sort"
)

// tptl. passed, but slow (23.8%). todo 2 find faster solution
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	ids := make([]int, len(nums1))
	for i := range ids {
		ids[i] = i
	}

	sort.Slice(ids, func(i, j int) bool {
		return nums1[ids[i]] < nums1[ids[j]]
	})

	bestImprovement, oldIDx, newIDx := 0, -1, -1
	for i := range nums1 {
		num2 := nums2[i]

		l, r := 0, len(ids)-1
		for l < r {
			m := l + (r-l)/2
			num1 := nums1[ids[m]]
			if num1 >= num2 {
				r = m
			} else {
				l = m + 1
			}
		}

		for _, idx := range [2]int{l, l - 1} {
			if idx >= 0 {
				diffBefore, diffAfter := abs(nums1[i]-num2), abs(nums1[ids[idx]]-num2)
				if diffBefore > diffAfter {
					improvement := diffBefore - diffAfter
					if improvement > bestImprovement {
						bestImprovement, newIDx, oldIDx = improvement, ids[idx], i
					}
				}
			}
		}
	}

	if newIDx != -1 {
		nums1[oldIDx] = nums1[newIDx]
	}

	res := 0
	for i := range nums1 {
		res += abs(nums1[i] - nums2[i])
		res %= 1_000_000_007
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
