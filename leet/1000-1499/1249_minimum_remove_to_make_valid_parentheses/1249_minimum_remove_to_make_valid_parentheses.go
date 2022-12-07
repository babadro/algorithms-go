package _1249_minimum_remove_to_make_valid_parentheses

// passed, but not so fast (and I don't know why)
func minRemoveToMakeValid3(s string) string {
	res := make([]byte, 0, len(s))
	opens := 0

	for i := range s {
		ch := s[i]
		if ch == ')' {
			if opens == 0 {
				continue
			}

			opens--
		} else if ch == '(' {
			opens++
		}

		res = append(res, ch)
	}

	if opens == 0 {
		return string(res)
	}

	// filter in place
	j := len(res) - 1
	for i := len(res) - 1; i >= 0; i-- {
		if opens > 0 && res[i] == '(' { // redundant opening parenthesis
			opens--
			continue
		}

		res[j] = res[i]
		j--
	}

	return string(res[j+1:])
}

// passed not mine. extremely fast (I don't know why)
func minRemoveToMakeValid(s string) string {
	res, opens := []byte{}, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			opens++
		} else if s[i] == ')' {
			if opens == 0 {
				continue
			}

			opens--
		}

		res = append(res, s[i])
	}

	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == '(' && opens > 0 {
			opens--
			res = append(res[:i], res[i+1:]...)
		}
	}

	return string(res)
}

// passed, fast, but too long
func minRemoveToMakeValid2(s string) string {
	b := []byte(s)
	var stack []int
	for i, ch := range b {
		if ch != '(' && ch != ')' {
			continue
		}

		if len(stack) > 0 {
			lastIDx := len(stack) - 1
			lastChar := b[stack[lastIDx]]

			if lastChar == '(' && ch == ')' {
				stack = stack[:lastIDx]
				continue
			}
		}

		stack = append(stack, i)
	}

	if len(stack) == 0 {
		return s
	}

	// mark invalid parentheses
	for _, idx := range stack {
		b[idx] = '_'
	}

	// filter in place
	i := 0
	for _, ch := range b {
		if ch == '_' {
			continue
		}

		b[i] = ch
		i++
	}

	return string(b[:i])
}
