package _507_perfect_number

// best solution. tptl
func checkPerfectNumber2(num int) bool {
	if num <= 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}

	return sum == num
}

// passed
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for d1, d2 := 2, num; d1 < d2; d1++ {
		if num%d1 == 0 {
			d2 = num / d1
			sum += d1 + d2
			if sum > num {
				return false
			}
		}
	}

	return sum == num
}
