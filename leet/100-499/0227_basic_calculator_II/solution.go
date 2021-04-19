package _227_basic_calculator_II

// best solution
// todo 2 understand
func calculate2(s string) int {
	number := 0
	index := 0
	toAdd := 1
	result := 0
	multiSign := true
	s = s + "+0"
	for index < len(s) {
		if s[index] <= '9' && s[index] >= '0' {
			number *= 10
			number += int(s[index] - '0')
		} else if s[index] == '+' || s[index] == '-' || s[index] == '/' || s[index] == '*' {
			if multiSign {
				toAdd *= number
			} else {
				toAdd /= number
			}
			if s[index] == '+' || s[index] == '-' {
				result += toAdd
				number = 0
				multiSign = true
				toAdd = 1
				if s[index] == '-' {
					toAdd = -1
				}
			} else {
				if s[index] == '/' {
					multiSign = false
				} else {
					multiSign = true
				}
				number = 0
			}
		}
		index++
	}
	return result
}
