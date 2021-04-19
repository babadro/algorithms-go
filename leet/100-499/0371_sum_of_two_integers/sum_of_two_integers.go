package _371_sum_of_two_integers

// TODO 2 need to understand. Good description here:
// https://www.geeksforgeeks.org/add-two-numbers-without-using-arithmetic-operators/
func getSum(a int, b int) int {
	for a != 0 {
		a, b = (a&b)<<1, a^b
	}
	return b
}

func getSum1(a int, b int) int {
	res, carry, sum, bit, aIn, bIn := int(0), int(0), int(0), int(0), int(0), int(0)
	for i := 0; i < 64; i++ {
		bit = int(1 << i)
		if a&bit == bit {
			aIn = 1
		} else {
			aIn = 0
		}
		if b&bit == bit {
			bIn = 1
		} else {
			bIn = 0
		}
		sum, carry = adder(aIn, bIn, carry)
		res |= sum << i
	}
	return res
}

func halfAdder(aIn, bIn int) (sumOut, carryOut int) {
	return aIn ^ bIn, aIn & bIn
}

func adder(aIn, bIn, carryIn int) (sumOut, carryOut int) {
	sum, carry1 := halfAdder(aIn, bIn)
	sumFinal, carry2 := halfAdder(sum, carryIn)
	carryFinal := carry1 | carry2
	return sumFinal, carryFinal
}
