package _2116_check_if_a_parentheses_string_can_be_valid

func canBeValid(s string, locked string) bool {
	var stack []int

	for i := range s {
		lock := locked[i] == '1'

		if !lock {
			stack = append(stack, i)
			continue
		}

		ch := s[i]

		if ch == '(' {
			stack = append(stack, i)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		last := len(stack) - 1

		prevIDx := stack[last]

		prevLock := locked[prevIDx] == '1'
		if !prevLock {
			stack = stack[:last]
			continue
		}

		if s[prevIDx] == '(' {
			stack = stack[:last]
			continue
		}

		return false
	}

	for i := range stack {
		if locked[stack[i]] == '1' {
			return false
		}
	}

	return len(stack)%2 == 0
}
