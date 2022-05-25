package _075_sort_colors

// tptl. passed. medium
func sortColors(nums []int) {
	low, high := 0, len(nums)-1
	for i := 0; i <= high; {
		if nums[i] == 0 {
			nums[i], nums[low] = nums[low], nums[i]
			i++
			low++
		} else if nums[i] == 1 {
			i++
		} else { // nums[i] == 2
			nums[i], nums[high] = nums[high], nums[i]
			high--
		}
	}
}
