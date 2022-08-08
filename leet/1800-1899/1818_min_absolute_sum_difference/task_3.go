package _1818_min_absolute_sum_difference

import (
	"sort"
)

// tptl. passed. fast. todo 3 can be improved further maybe
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	copyNums1 := make([]int, len(nums1))
	copy(copyNums1, nums1)

	sort.Ints(copyNums1)

	bestImprovement, targetIDx, num := 0, -1, 0
	for i := range nums1 {
		num2 := nums2[i]

		l, r := 0, len(nums1)-1
		for l < r {
			m := l + (r-l)/2
			num1 := copyNums1[m]
			if num1 >= num2 {
				r = m
			} else {
				l = m + 1
			}
		}

		for _, idx := range [2]int{l, l - 1} {
			if idx >= 0 {
				diffBefore, diffAfter := abs(nums1[i]-num2), abs(copyNums1[idx]-num2)
				if diffBefore > diffAfter {
					improvement := diffBefore - diffAfter
					if improvement > bestImprovement {
						bestImprovement, num, targetIDx = improvement, copyNums1[idx], i
					}
				}
			}
		}
	}

	if targetIDx != -1 {
		nums1[targetIDx] = num
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
