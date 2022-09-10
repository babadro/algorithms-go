package _136_single_number

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}

	return res
}
