package _260_single_number_III

// passed. tptl
func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	firstDiffBit := 1
	for firstDiffBit&xor == 0 {
		firstDiffBit <<= 1
	}

	num1, num2 := 0, 0
	for _, num := range nums {
		if num&firstDiffBit == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}
