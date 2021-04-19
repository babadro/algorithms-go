package _022_generate_parentheses

// TODO 2 implement approach 1 and 2 and 3 from solutions
func generateParenthesis2(n int) []string {
	var combinations []string
	generateAll(make([]byte, 2*n), 0, combinations)
	return combinations
}

// TODO 2 implement recursive generator
func generateAll(current []byte, pos int, result []string) {
	if pos == len(current) {
		if valid(current) {
			result = append(result, string(current))
		} else {
			current[pos] = '('
			generateAll(current, pos+1, result)
			current[pos] = ')'
			generateAll(current, pos+1, result)
		}
	}
}

func valid(current []byte) bool {
	balance := 0
	for _, c := range current {
		if c == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}
