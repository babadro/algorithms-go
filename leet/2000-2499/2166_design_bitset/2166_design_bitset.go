package _2166_design_bitset

// tptl. passed
type Bitset struct {
	strBuf, flipStrBuf []byte
	dict, flipDict     map[int]bool
}

func Constructor(size int) Bitset {
	str, flipStr := make([]byte, size), make([]byte, size)
	dict, flipDict := make(map[int]bool), make(map[int]bool)

	for i := range str {
		str[i] = '0'
		flipStr[i] = '1'

		flipDict[i] = true
	}

	return Bitset{
		strBuf:     str,
		flipStrBuf: flipStr,
		dict:       dict,
		flipDict:   flipDict,
	}
}

func (this *Bitset) Fix(idx int) {
	if _, ok := this.dict[idx]; !ok {
		this.dict[idx] = true
		this.strBuf[idx] = '1'

		delete(this.flipDict, idx)
		this.flipStrBuf[idx] = '0'
	}
}

func (this *Bitset) Unfix(idx int) {
	if _, ok := this.dict[idx]; ok {
		delete(this.dict, idx)
		this.strBuf[idx] = '0'

		this.flipDict[idx] = true
		this.flipStrBuf[idx] = '1'
	}
}

func (this *Bitset) Flip() {
	this.strBuf, this.flipStrBuf = this.flipStrBuf, this.strBuf
	this.dict, this.flipDict = this.flipDict, this.dict
}

func (this *Bitset) All() bool {
	return len(this.dict) == len(this.strBuf)
}

func (this *Bitset) One() bool {
	return len(this.dict) > 0
}

func (this *Bitset) Count() int {
	return len(this.dict)
}

func (this *Bitset) ToString() string {
	return string(this.strBuf)
}
