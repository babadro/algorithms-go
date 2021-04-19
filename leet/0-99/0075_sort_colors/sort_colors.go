package _075_sort_colors

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	j, k := 0, len(nums)-1
	i := 0
	for i <= k {
		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[k] = nums[k], nums[i]
			k--
		} else {
			i++
		}
	}
}
