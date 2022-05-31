package _1009_complement_of_base_10_integer

// tptl. passed
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	bitCount := 0
	for num := n; num > 0; num >>= 1 {
		bitCount++
	}

	allBitsSet := (1 << bitCount) - 1

	return n ^ allBitsSet
}
