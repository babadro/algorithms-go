package _191_number_of_1_bits

func hammingWeight(num uint32) int {
	x := uint64(num)
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func hammingWeight2(num uint32) int {
	sum := 0
	for num != 0 {
		sum++
		num &= num - 1
	}
	return sum
}

// best solution
func hammingWeight3(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
