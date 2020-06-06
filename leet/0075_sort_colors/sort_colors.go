package _075_sort_colors

// TODO 1
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	j, k := 0, 0
	for i, v := range nums {
		if v != 0 {
			j = i
			break
		}
	}
	for i := len(nums) - 1; i > j; i-- {
		if nums[i] != 2 {
			k = i
			break
		}
	}
	i := j
	for i <= k {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		} else if nums[i] == 2 {
			nums[i], nums[k] = nums[k], nums[i]
			k--
		}
	}
}
