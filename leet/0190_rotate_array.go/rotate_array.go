package _190_rotate_array_go

func rotate(nums []int, k int) {
	l := len(nums)
	var tmp int
	for i := 0; i < k; i++ {
		tmp = nums[i]
		nums[i] = nums[l-k+i]
		tmp, nums[i+k] = nums[i+k], tmp
		tmp, nums[i+2*k] = nums[i+2*k], tmp
		nums[i+3*k] = tmp

	}
}
