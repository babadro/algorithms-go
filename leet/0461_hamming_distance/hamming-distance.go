package _461_hamming_distance

import "math/bits"

func hammingDistance(x int, y int) int {
	return bitCount(uint64(x ^ y))
}

func bitCount(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func hammingDistance2(x int, y int) int {
	d := 0
	for x > 0 || y > 0 {
		if x&1 != y&1 {
			d++
		}
		x >>= 1
		y >>= 1
	}
	return d
}

func hammingDistance3(x int, y int) int {
	return bits.OnesCount(uint(x) ^ uint(y))
}
