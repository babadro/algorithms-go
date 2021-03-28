package _34

func numDifferentIntegers(word string) int {
	m := make(map[string]bool)
	counter := 0
	n := len(word)
	digit := make([]byte, 0)
	for i := 0; i <= n; i++ {
		if i < n {
			char, nextIdx := word[i], i+1

			if char == '0' && (len(digit) > 0 || nextIdx == n || (word[nextIdx] < '0' || word[nextIdx] > '9')) ||
				(char > '0' && char <= '9') {
				digit = append(digit, char)
				continue
			}
		}

		if len(digit) > 0 {
			if !m[string(digit)] {
				counter++
				m[string(digit)] = true
			}

			digit = digit[:0]
		}
	}

	return counter
}
