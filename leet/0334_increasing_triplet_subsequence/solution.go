package _334_increasing_triplet_subsequence

import "math"

// 91% 100%
// best solution
// https://leetcode.com/problems/increasing-triplet-subsequence/discuss/822326/Go-O(N)-with-math-import-for-an-easy-max-value
func increasingTriplet(nums []int) bool {
	firstNum, secondNum := math.MaxInt64, math.MaxInt64

	for i := 0; i < len(nums); i++ {
		if nums[i] <= firstNum {
			firstNum = nums[i]
		} else if nums[i] <= secondNum {
			secondNum = nums[i]
		} else {
			return true
		}
	}

	return false
}
