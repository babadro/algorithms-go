package _1_maximum_sum_subarray_of_size_k

import "github.com/babadro/algorithms-go/utils"

func findMaxSumSubArray(k int, arr []int) int {
	sum, start, max := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i >= k-1 {
			max = utils.Max(max, sum)
			sum -= arr[start]
			start++
		}
	}

	return max
}
