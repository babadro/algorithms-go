package _2_smallest_subarray_with_a_greater_sum

import (
	"math"

	"github.com/babadro/algorithms-go/utils"
)

// tptl
// Given an array of positive numbers and a positive number ‘S,’
// find the length of the smallest contiguous subarray whose sum is greater than or equal to ‘S’.
// Return 0 if no such subarray exists.
func findMinSubArray(s int, arr []int) int {
	minLen, sum, start := math.MaxInt64, 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		for sum >= s {
			minLen = utils.Min(minLen, i-start+1)
			sum -= arr[start]
			start++
		}
	}

	if minLen == math.MaxInt64 {
		return 0
	}

	return minLen
}
