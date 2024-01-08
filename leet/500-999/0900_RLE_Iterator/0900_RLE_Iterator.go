package _900_RLE_Iterator

type RLEIterator struct {
	encoding []int
	i        int
}

// #bnsrg passed medium
func Constructor(encoding []int) RLEIterator {
	return RLEIterator{encoding: encoding}
}

func (this *RLEIterator) Next(n int) int {
	for {
		if this.i >= len(this.encoding) {
			return -1
		}

		val := this.encoding[this.i]

		if val == 0 {
			this.i += 2
			continue
		}

		if n <= val {
			this.encoding[this.i] = val - n

			return this.encoding[this.i+1]
		}

		n -= val
		this.i += 2
	}
}
