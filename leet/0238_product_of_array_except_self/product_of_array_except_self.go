package _238_product_of_array_except_self

// Approach 2, O(n), O(1)
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	left[0] = 1

	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		left[i] = left[i] * right
		right *= nums[i]
	}

	return left
}

// Approach 1, O(n), O(n)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = 1, 1

	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = left[i] * right[i]
	}

	return res
}
