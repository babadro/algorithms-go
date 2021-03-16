package _1793_maximum_score_of_a_good_subarray

import "github.com/babadro/algorithms-go/utils"

// best solution. tptl. passed. hard.
func maximumScore2(nums []int, k int) int {
	i, j, minNum, maxScore := k, k, nums[k], nums[k]

	n := len(nums)
	for i > 0 || j < n-1 {
		if i > 0 && j < n-1 {
			if nums[j+1] > nums[i-1] {
				j++
			} else {
				i--
			}
		} else if i == 0 && j < n-1 {
			j++
		} else if j == n-1 && i > 0 {
			i--
		}

		minNum = utils.Min3(nums[i], nums[j], minNum)
		score := minNum * (j - i + 1)

		maxScore = utils.Max(maxScore, score)
	}

	return maxScore
}
