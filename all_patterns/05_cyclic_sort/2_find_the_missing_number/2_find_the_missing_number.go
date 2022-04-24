package __find_the_missing_number

// leetcode 286 (there is a simpler and trickier solution - look at leetcode task 286 implementation)
func findMissingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] < len(nums) && nums[i] != i {
			tmp := nums[i]
			nums[i] = nums[nums[i]]
			nums[tmp] = tmp
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}
