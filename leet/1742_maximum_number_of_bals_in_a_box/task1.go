package _1742_maximum_number_of_bals_in_a_box

// contest. passed
func countBalls(lowLimit int, highLimit int) int {
	m := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		res := digitsSum(i)
		m[res]++
	}

	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	return max
}

func digitsSum(num int) int {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit
		num = num / 10
	}

	return sum
}
