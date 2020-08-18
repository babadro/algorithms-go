package _227_basic_calculator_II

// TODO 1
func calculate(s string) int {
	n := len(s)
	left, right, add := 0, 0, false
	for i := 0; i < n; i++ {
		i, add = parsePlusOrMinus(i, s)
		if i == n {
			return left
		}
		i, right = calcOperand(i, s)
		left = evalAddOrSub(left, right, add)
	}

	return left
}

func calcOperand(i int, s string) (int, int) {
	n := len(s)
	left, right, mult := 0, 0, false
	for ; i < len(s) && !(s[i] == '+' || s[i] == '-'); i++ {
		i, left = parseNum(i, s)
		if i == n || s[i] == '+' || s[i] == '-' {
			return i, left
		}
		i, mult = parseDivOrMult(i, s)
		if i == n {
			return i, left
		}
		i, right = parseNum(i, s)
		left = evalMultOrDiv(left, right, mult)
	}

	return i, left
}

func parseNum(i int, s string) (int, int) {
	for s[i] == ' ' {
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
	for s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return i, false
	}
	if s[i] == '-' {
		return i, false
	}
	return i, true
}

func parseDivOrMult(i int, s string) (int, bool) {
	for s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return i, false
	}
	if s[i] == '/' {
		return i, false
	}
	return i, true
}

func evalAddOrSub(left, right int, add bool) int {
	if add {
		return left + right
	}
	return left - right
}

func evalMultOrDiv(left, right int, mult bool) int {
	if mult {
		return left * right
	}
	return left / right
}