package _560_subarray_sum_equals_k

// tptl. passed
func subarraySum(nums []int, k int) int {
	sumFreq := map[int]int{0: 1}
	res, sum := 0, 0
	for _, num := range nums {
		sum += num
		res += sumFreq[sum-k]
		sumFreq[sum]++
	}

	return res
}
