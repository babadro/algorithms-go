package _190_reverse_bits

// passed. tptl. easy
func reverseBits(num uint32) uint32 {
	var res uint32
	bit := uint32(1 << 31)
	for i := 0; i < 32; i++ {
		if 1<<i&num != 0 {
			res |= bit >> i
		}
	}

	return res
}
