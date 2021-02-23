package _1417_reformat_the_string

// tptl. passed
func reformat(s string) string {
	n := len(s)
	i, j := 0, n-1
	b := make([]byte, n)

	for k := range s {
		char := s[k]
		if digit(char) {
			b[i] = char
			i++
		} else {
			b[j] = char
			j--
		}
	}

	if diff := n - 2*i; diff < -1 || diff > 1 {
		return ""
	}

	j++
	i--

	remainder := n/2 - i
	if remainder != 0 && remainder != 1 {
		return ""
	}

	for k, char := range b {
		if k%2 == remainder {
			if !digit(char) {
				b[k], b[j] = b[j], b[k]
				j++
			}
		} else if digit(char) {
			b[k], b[j] = b[j], b[k]
			j++
		}
	}

	return string(b)
}

func digit(char byte) bool {
	return char >= '0' && char <= '9'
}
