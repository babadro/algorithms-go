package _470_implemend_rand10_using_rand7

import "math/rand"

// # bnsrg medium passed
// todo 2 there are very short solution
// todo 3 there are optimization opportunity to minimize number of rand7() api calls
func rand10() int {
	for {
		base7Num := [2]int{rand7() - 1, rand7() - 1}

		base10Num := 0
		// convert to base10Num
		for _, digit := range base7Num {
			base10Num *= 7
			base10Num += digit % 7
		}

		base10Num += 1

		if base10Num > 40 {
			continue
		}

		return base10Num % 10
	}
}

func rand7() int {
	return rand.Intn(7) + 1
}
