package _075_sort_colors

// tptl. passed. medium
func sortColors(nums []int) {
	firstOne, firstTwo := 0, len(nums)-1
	for i := 0; i <= firstTwo; {
		if nums[i] == 0 {
			nums[i], nums[firstOne] = nums[firstOne], nums[i]
			i++
			firstOne++
		} else if nums[i] == 1 {
			i++
		} else { // nums[i] == 2
			nums[i], nums[firstTwo] = nums[firstTwo], nums[i]
			firstTwo--
		}
	}
}
