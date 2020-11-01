package _1021_remove_outermost_parentheses

// tptl
// passed.
func removeOuterParentheses(S string) string {
	res := make([]byte, 0, len(S))

	counter, lastIdx := 0, 0
	for i, v := range S {
		if v == '(' {
			counter++
		} else {
			counter--
		}
		if counter == 0 {
			res = append(res, S[lastIdx+1:i]...)
			lastIdx = i + 1
		}
	}

	return string(res)
}
