package _0_introduction

// find the average of all subarrays of k contiguous elements in the given array.
func avgOfSubarrayOfSizeK(k int, arr []int) []float64 {
	res := make([]float64, len(arr)-k+1)

	sum, start := 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i >= k-1 {
			res[start] = float64(sum) / float64(k)
			sum -= arr[start]
			start++
		}
	}

	return res
}
