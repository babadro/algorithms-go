package _020_valid_parentheses

func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	stack := make([]rune, length/2)
	idx := 0

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			if idx == len(stack) {
				return false
			}
			stack[idx] = v
			idx++
		} else {
			if idx == 0 {
				return false
			}
			if stack[idx-1] != paired(v) {
				return false
			}
			idx--
		}
	}
	return idx == 0
}

func paired(close rune) rune {
	switch close {
	case '}':
		return '{'
	case ')':
		return '('
	default:
		return '['
	}
}
