package _043_multiply_strings

func multiply(num1 string, num2 string) string {
	// some checks...

	l1, l2 := len(num1), len(num2)
	s, l := num1, num2
	if l2 < l1 {
		s, l = num2, num1
	}

	ld := make([]int, len(l))
	for j := 0; j < len(ld); j++ {
		ld[j] = int(l[j] - '0')
	}

	result := make([]int, (l1+l2)*2)
	lastDigit := 0

	for j, start := len(s)-1, 0; j >= 0; j, start = j-1, start+1 {
		lastDigit = multiply(ld)
	}
}

func mult(a []int, n int, result []int, start int) int {
	carryOver := 0

	for j := len(a) - 1; j >= 0; j, start = j-1, start+1 {
		res := (a[j] * n) + carryOver + result[start]
		if res > 9 {
			carryOver = res / 10
			result[start] = res - (carryOver * 10)
		} else {
			carryOver = 0
			result[start] = res
		}
	}

	for ; carryOver > 0; start++ {
		res := carryOver + result[start]
		if res > 9 {
			carryOver = res / 10
			result[start] = res - (carryOver * 10)
		} else {
			carryOver = 0
			result[start] = res
		}
	}

	return start - 1
}
