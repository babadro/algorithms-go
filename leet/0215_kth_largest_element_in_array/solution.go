package _215_kth_largest_element_in_array

// https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/538760/Go-quick-select
// TODO 1 Need to understand
func findKthLargest2(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	pivot := nums[0]
	p1, p2 := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= pivot {
			p2++
		} else {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
			p2++
		}
	}
	nums[0], nums[p1-1] = nums[p1-1], nums[0]
	if p1 == len(nums)-k+1 {
		return pivot
	} else if p1 < len(nums)-k+1 {
		return findKthLargest2(nums[p1:], k)
	} else {
		return findKthLargest2(nums[:p1-1], p1+k-len(nums)-1)
	}
}
