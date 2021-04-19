package _260_single_number_III

// passed. tptl
func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	var bit int
	for i := 0; i < 64; i++ {
		if xor&(1<<i) != 0 {
			bit = 1 << i
			break
		}
	}

	var x, y int
	for _, num := range nums {
		if num&bit != 0 {
			x ^= num
		} else {
			y ^= num
		}
	}

	return []int{x, y}
}
