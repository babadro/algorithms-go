package _2166_design_bitset

import "math"

const leftBit = uint64(1) << 63

type Bitset struct {
	arr          []uint64
	size         int
	lastItemMask uint64
	strBuf       []byte
	bitCount     int
}

func Constructor(size int) Bitset {
	arrLen := size / 64
	remainder := size % 64
	var lastItemMask uint64
	if remainder != 0 {
		for i := 0; i < remainder; i++ {
			lastItemMask |= leftBit >> i
		}

		arrLen++
	} else {
		lastItemMask = math.MaxUint64
	}

	strBuf := make([]byte, size)
	for i := range strBuf {
		strBuf[i] = '0'
	}

	return Bitset{
		arr:          make([]uint64, arrLen),
		size:         size,
		lastItemMask: lastItemMask,
		strBuf:       strBuf,
	}
}

func (this *Bitset) Fix(idx int) {
	bucketIDx, shift := getIDs(idx)
	bit := leftBit >> shift

	if this.strBuf[idx] == '0' {
		this.bitCount++
	}

	this.strBuf[idx] = '1'
	this.arr[bucketIDx] |= bit
}

func (this *Bitset) Unfix(idx int) {
	if this.strBuf[idx] == '1' {
		this.bitCount--
	}

	this.strBuf[idx] = '0'

	bucketIDx, shift := getIDs(idx)
	mask := math.MaxUint64 ^ (leftBit >> shift)

	this.arr[bucketIDx] &= mask
}

func (this *Bitset) Flip() {
	this.bitCount = this.size - this.bitCount

	for i, ch := range this.strBuf {
		if ch == '0' {
			this.strBuf[i] = '1'
		} else {
			this.strBuf[i] = '0'
		}
	}

	for i, num := range this.arr {
		if i == len(this.arr)-1 {
			this.arr[i] = num ^ this.lastItemMask
		} else {
			this.arr[i] = num ^ math.MaxUint64
		}
	}
}

func (this *Bitset) All() bool {
	for i, num := range this.arr {
		allOnesMask := uint64(math.MaxUint64)
		if i == len(this.arr)-1 {
			allOnesMask = this.lastItemMask
		}

		if num != allOnesMask {
			return false
		}
	}

	return true
}

func (this *Bitset) One() bool {
	for _, num := range this.arr {
		if num != 0 {
			return true
		}
	}

	return false
}

func (this *Bitset) Count() int {
	return this.bitCount
}

func (this *Bitset) ToString() string {
	return string(this.strBuf)
}

func getIDs(idx int) (int, int) {
	bucketIDx := idx / 64
	remainder := idx % 64

	return bucketIDx, remainder
}
