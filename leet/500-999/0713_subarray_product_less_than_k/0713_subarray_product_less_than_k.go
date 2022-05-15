package _713_subarray_product_less_than_k

// passed. tptl. medium #sliding window
func numSubarrayProductLessThanK(nums []int, k int) int {
	product := 1
	left := 0

	res := 0
	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		for product >= k && left < len(nums) {
			product /= nums[left]
			left++
		}

		if left <= right {
			res += right - left + 1
		}
	}

	return res
}
