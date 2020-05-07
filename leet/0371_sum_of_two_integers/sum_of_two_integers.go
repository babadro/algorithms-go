package _371_sum_of_two_integers

// TODO 1
func getSum(a int8, b int8) int8 {
	res, carry, sum, bit, aIn, bIn := int8(0), int8(0), int8(0), int8(0), int8(0), int8(0)
	for i := 0; i < 8; i++ {
		bit = int8(1 << i)
		if a&bit > 0 {
			aIn = 1
		} else {
			aIn = 0
		}
		if b&bit > 0 {
			bIn = 1
		} else {
			bIn = 0
		}
		sum, carry = adder(aIn, bIn, carry)
		res |= sum << i
	}
	return res
}

func halfAdder(aIn, bIn int8) (sumOut, carryOut int8) {
	return aIn ^ bIn, aIn & bIn
}

func adder(aIn, bIn, carryIn int8) (sumOut, carryOut int8) {
	sum, carry1 := halfAdder(aIn, bIn)
	sumFinal, carry2 := halfAdder(sum, carryIn)
	carryFinal := carry1 | carry2
	return sumFinal, carryFinal
}
