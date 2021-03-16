package _1793_maximum_score_of_a_good_subarray

import "math"

// todo 1 doesnt' work
func maximumScore(nums []int, k int) int {
	leftIdx, leftMin, leftMaxScore := k, math.MaxInt64, 0
	for i := k; i >= 0; i-- {
		if num := nums[i]; num < leftMin {
			leftMin = num
		}

		length := k - i + 1

		if score := leftMin * length; score > leftMaxScore {
			leftIdx, leftMaxScore = i, score
		}
	}

	rightIdx, rightMin, rightMaxScore := k, math.MaxInt64, 0
	for i := k; i < len(nums); i++ {
		if num := nums[i]; num < rightMin {
			rightMin = num
		}

		length := i - k + 1

		if score := rightMin * length; score > rightMaxScore {
			rightIdx, rightMaxScore = i, score
		}
	}

	min := math.MaxInt64
	for i := leftIdx; i <= rightIdx; i++ {
		if num := nums[i]; num < min {
			min = num
		}
	}

	totalScore := (rightIdx - leftIdx + 1) * min

	max := leftMaxScore
	if rightMaxScore > leftMaxScore {
		max = rightMaxScore
	}

	if totalScore > max {
		max = totalScore
	}

	return max
}

// Time limit exceeded. But, probably, algo is correct.
func maximumScoreBruteForce(nums []int, k int) int {
	leftMin := math.MaxInt64
	for i := k; i >= 0; i-- {
		if num := nums[i]; num < leftMin {
			leftMin = num
		}

		nums[i] = leftMin
	}

	rightMin := math.MaxInt64
	for i := k; i < len(nums); i++ {
		if num := nums[i]; num < rightMin {
			rightMin = num
		}

		nums[i] = rightMin
	}

	maxScore, mid := 0, nums[k]
	for i := k; i >= 0; i-- {
		for j := k; j < len(nums); j++ {
			left, right := nums[i], nums[j]

			minNum := min(left, min(right, mid))

			length := j - i + 1

			if score := minNum * length; score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
