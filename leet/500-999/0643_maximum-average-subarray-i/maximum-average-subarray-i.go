package _643_maximum_average_subarray_i

// passed. easy
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	max := sum
	for i, j := 0, k; j < len(nums); j, i = j+1, i+1 {
		sum += nums[j]
		sum -= nums[i]

		if sum > max {
			max = sum
		}
	}

	return float64(max) / float64(k)
}
