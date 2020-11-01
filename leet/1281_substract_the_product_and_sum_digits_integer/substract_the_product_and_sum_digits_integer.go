package _1281_substract_the_product_and_sum_digits_integer

// tptl. passed
func subtractProductAndSum(n int) int {
	if n == 0 {
		return 0
	}

	digits := make([]int, 0, 6)

	divider := 10
	for n > 0 {
		digit := n % divider
		digits = append(digits, digit)
		n = n / divider
	}

	product, sum := digits[0], digits[0]
	for _, v := range digits[1:] {
		product *= v
		sum += v
	}

	return product - sum
}
