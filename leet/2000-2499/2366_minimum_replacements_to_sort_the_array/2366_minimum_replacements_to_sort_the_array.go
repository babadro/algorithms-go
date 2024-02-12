package _2366_minimum_replacements_to_sort_the_array

// #bnsrg hard passed
func minimumReplacement(nums []int) int64 {
	ops := 0

	for i := len(nums) - 2; i >= 0; i-- {
		length := nums[i] / nums[i+1]
		if nums[i]%nums[i+1] > 0 {
			length += 1
		}

		nums[i] /= length

		ops += length - 1
	}

	return int64(ops)
}
