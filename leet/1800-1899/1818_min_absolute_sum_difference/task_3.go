package _1818_min_absolute_sum_difference

import "github.com/babadro/algorithms-go/utils"

// todo 1 doesn't work
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
	distances := make([]int, n)

	sum := 0
	maxDistIdx := -1
	for i := 0; i < n; i++ {
		dist := utils.Abs(nums1[i] - nums2[i])
		distances[i] = dist
		sum += dist

		if maxDistIdx == -1 || dist > distances[maxDistIdx] {
			maxDistIdx = i
		}
	}

	minDiffSum := sum
	for i := 0; i < n; i++ {
		curr := sum - distances[maxDistIdx] + utils.Abs(nums1[i]-nums2[maxDistIdx])

		if curr < minDiffSum {
			minDiffSum = curr
		}
	}

	return minDiffSum % 1_000_000_007
}
