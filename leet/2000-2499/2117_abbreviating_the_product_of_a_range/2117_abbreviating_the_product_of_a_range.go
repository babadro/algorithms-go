package _2117_abbreviating_the_product_of_a_range

import "math"

// todo 1
func abbreviateProduct(left int, right int) string {
	count2, count5 := 0, 0
	last5digits := 1
	var power10 float64
	for num := left; num <= right; num++ {
		// calc trailing zeroes
		count2 += countMultipliers(num, 2)
		count5 += countMultipliers(num, 5)

		// calc last 5 digits
		last5digits *= num
		last5digits %= 100000

		// calc first 5 digits
		power10 += math.Log10(float64(num))
	}

	zeroesCount := min(count2, count5)

	fractionPart := power10 - float64(int(power10))
	first5Digits := int(math.Floor(math.Pow(10, fractionPart+4)))

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func countMultipliers(num, multiplier int) int {
	count := 0
	for num%multiplier == 0 {
		count++
	}

	return count
}
