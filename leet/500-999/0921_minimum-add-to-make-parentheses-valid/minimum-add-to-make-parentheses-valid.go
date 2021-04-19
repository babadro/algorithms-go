package _921_minimum_add_to_make_parentheses_valid

// passed. medium. tptl
func minAddToMakeValid(S string) int {
	stack := make([]byte, 0)

	for i := range S {
		b, last := S[i], len(stack)-1
		if len(stack) > 0 && stack[last] == '(' && b == ')' {
			stack = stack[:last]
		} else {
			stack = append(stack, b)
		}
	}

	return len(stack)
}
