package _728_self_dividing_numbers

// easy. passed.
func selfDividingNumbers(left int, right int) []int {
	var res []int

Loop:
	for i := left; i <= right; i++ {
		num := i
		for num > 0 {
			digit := num % 10
			if digit == 0 || i%digit != 0 {
				continue Loop
			}

			num = num / 10
		}

		res = append(res, i)
	}

	return res
}
