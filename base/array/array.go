package array

func filterInPlacePreserveOrder(nums []int, f func(num int) bool) []int {
	i := 0
	for _, num := range nums {
		if f(num) {
			nums[i] = num
			i++
		}
	}

	return nums[:i]
}

func filterInPlaceNotPreserveOrder(nums []int, f func(num int) bool) []int {
	for i := 0; i < len(nums); {
		if f(nums[i]) {
			i++
			continue
		}

		last := len(nums) - 1
		nums[i], nums[last] = nums[last], nums[i]

		nums = nums[:last]
	}

	return nums
}
