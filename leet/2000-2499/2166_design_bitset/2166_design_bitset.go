package _2166_design_bitset

import "math"

const leftBit = uint64(1) << 63

type Bitset struct {
	arr, flipArr           []uint64
	size                   int
	lastItemMask           uint64
	strBuf, flipStrBuf     []byte
	bitCount, flipBitCount int
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

	strBuf, flipStrBUf := make([]byte, size), make([]byte, size)
	for i := range strBuf {
		strBuf[i] = '0'
		flipStrBUf[i] = '1'
	}

	arr, flipArr := make([]uint64, arrLen), make([]uint64, arrLen)
	for i := range flipArr {
		flipArr[i] = math.MaxUint64
	}

	return Bitset{
		arr:          arr,
		flipArr:      flipArr,
		size:         size,
		lastItemMask: lastItemMask,
		strBuf:       strBuf,
		flipStrBuf:   flipStrBUf,
		bitCount:     0,
		flipBitCount: size,
	}
}

func (this *Bitset) Fix(idx int) {
	this.fix(idx, false)
	this.unfix(idx, true)
}

func (this *Bitset) fix(idx int, flip bool) {
	bucketIDx, shift := getIDs(idx)
	bit := leftBit >> shift

	bitCount, arr, str := &this.bitCount, this.arr, this.strBuf
	if flip {
		bitCount, arr, str = &this.flipBitCount, this.flipArr, this.flipStrBuf
	}

	if str[idx] == '0' {
		*bitCount++
	}

	str[idx] = '1'
	arr[bucketIDx] |= bit
}

func (this *Bitset) Unfix(idx int) {
	this.unfix(idx, false)
	this.fix(idx, true)
}

func (this *Bitset) unfix(idx int, flip bool) {
	bitCount, arr, str := &this.bitCount, this.arr, this.strBuf
	if flip {
		bitCount, arr, str = &this.flipBitCount, this.flipArr, this.flipStrBuf
	}

	if str[idx] == '1' {
		*bitCount--
	}

	str[idx] = '0'

	bucketIDx, shift := getIDs(idx)
	mask := math.MaxUint64 ^ (leftBit >> shift)

	arr[bucketIDx] &= mask
}

func (this *Bitset) Flip() {
	this.arr, this.flipArr = this.flipArr, this.arr
	this.strBuf, this.flipStrBuf = this.flipStrBuf, this.strBuf
	this.bitCount, this.flipBitCount = this.flipBitCount, this.bitCount
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
