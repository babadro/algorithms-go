package _371_sum_of_two_integers

// TODO 1
func getSum(a int8, b int8) int8 {
	res, carry, bit := int8(0), int8(0), int8(0)
	for i := 0; i < 8; i++ {
		bit = int8(1 << i)
		bitA, bitB := a&bit > 0, b&bit > 0
		if bitA && bitB {
			if carry == 1 {
				res |= bit
			} else {
				carry = 1
			}
		} else if (bitA || bitB) && carry == 0 {
			res |= bit
		} else if carry == 1 {
			res |= bit
			carry = 0
		}
	}
	return res
}
