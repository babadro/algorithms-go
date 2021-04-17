package _807_evaluate_the_bracket_pairs_of_a_string

// passed. todo 1. can i do it without key array?
func evaluate(s string, knowledge [][]string) string {
	dict := make(map[string][]byte, len(knowledge))

	for _, kvp := range knowledge {
		k, v := kvp[0], kvp[1]

		dict[k] = []byte(v)
	}

	b := make([]byte, 0, len(s))

	key, keyParsing := make([]byte, 0), false
	for i := range s {
		char := s[i]
		if !keyParsing {
			if char != '(' {
				b = append(b, char)
			} else {
				keyParsing = true
			}
		} else {
			if char == ')' {
				keyParsing = false
				value, ok := dict[string(key)]
				if !ok {
					b = append(b, '?')
				} else {
					b = append(b, value...)
				}
				key = key[:0]
			} else {
				key = append(key, char)
			}
		}
	}

	return string(b)
}
