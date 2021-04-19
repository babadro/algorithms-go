package _198_house_robber

// tptl. passed. array.
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	nums[1] = max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}
	return nums[l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
