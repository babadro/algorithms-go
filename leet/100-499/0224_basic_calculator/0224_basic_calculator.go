package _224_basic_calculator

// tptl. passed. hard
func calculate(s string) int {
	res, _ := rec(s, 0)

	return res
}

func rec(s string, i int) (res, idx int) {
	minus, num := false, 0
	for i < len(s) {
		ch := s[i]
		if ch == ')' {
			return res, i + 1
		}

		if ch == ' ' || ch == '+' {
			i++
			continue
		}

		if ch == '-' {
			minus = true
			i++
			continue
		}

		if ch >= '0' && ch <= '9' {
			num, i = parseNum(s, i)
		} else if ch == '(' {
			num, i = rec(s, i+1)
		}

		if minus {
			res -= num
		} else {
			res += num
		}

		minus = false
	}

	return res, i
}

func parseNum(s string, i int) (num, idx int) {
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		num *= 10
		num += int(s[i] - '0')
	}

	return num, i
}
