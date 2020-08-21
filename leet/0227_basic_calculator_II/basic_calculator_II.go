package _227_basic_calculator_II

// 84% 94%
func calculate(s string) int {
	n := len(s)
	left, right, add := 0, 0, false
	for i := 0; i < n; {
		i, add = parsePlusOrMinus(i, s)
		if i == n {
			return left
		}
		i, right = calcOperand(i, s)

		if add {
			left += right
		} else {
			left -= right
		}
	}

	return left
}

func calcOperand(i int, s string) (int, int) {
	n := len(s)
	left, right, mult, ok := 0, 0, false, false
	i, left = parseNum(i, s)
	for i < n {
		i, mult, ok = parseMultOrDiv(i, s)
		if !ok {
			return i, left
		}
		i, right = parseNum(i, s)

		if mult {
			left *= right
		} else {
			left /= right
		}
	}

	return i, left
}

func parseNum(i int, s string) (int, int) {
	for i < len(s) && s[i] == ' ' {
		i++
	}
	res := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		num := s[i] - '0'
		res *= 10
		res += int(num)
	}

	return i, res
}

func parsePlusOrMinus(i int, s string) (int, bool) {
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return i, false
	}
	if s[i] == '-' {
		return i + 1, false
	}
	if s[i] == '+' {
		return i + 1, true
	}
	return i, true
}

func parseMultOrDiv(i int, s string) (int, bool, bool) {
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return i, false, false
	}
	if s[i] == '/' {
		return i + 1, false, true
	}
	if s[i] == '*' {
		return i + 1, true, true
	}

	return i, false, false
}
