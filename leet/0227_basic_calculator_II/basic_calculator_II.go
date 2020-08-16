package _227_basic_calculator_II

const add = 0
const sub = 1
const mult = 2
const div = 3

// TODO 1
func calculate(s string) int {

}

func calc(s string, multAndDivOnly bool) int {
	n := len(s)
	left := 0
	var operator int
	var right int
	for i := 0; i < n; i++ {
		if i > 0 {
			i, left = findOperand(i, s)
			if i == n {
				return left
			}

			i, operator = findOperator(i, s)
			if i == n {
				return left
			}
		}

		i, right = findOperand(i, s)
		left = eval(left, right, operator)
	}

	return left
}

func findOperand(i int, s string, multAndDivOnly bool) (int, int) {
	for s[i] == ' ' {
		i++
	}
	start := i
	res := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		num := s[i] - '0'
		res *= 10
		res += int(num)
	}

	if i < len(s) && !multAndDivOnly && (s[i] == '/' || s[i] == '*') {
		for ; i < len(s) && !(s[i] == ' ' || s[i] == '+' || s[i] == '-'); i++ {
		}
		return i, calc(s[start:i], true)
	}

	return i, res
}

func findOperator(i int, s string) (int, int) {
	return 0, 0
}

func eval(left, right, operator int) int {
	return 0
}
