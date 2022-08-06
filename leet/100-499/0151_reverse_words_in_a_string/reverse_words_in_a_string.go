package _151_reverse_words_in_a_string

// passed. medium. #string
func reverseWords(s string) string {
	b := []byte(s)

	i := 0
	for j := 0; j < len(b); j++ {
		if b[j] == ' ' && (i == 0 || b[i-1] == ' ') {
			continue
		}

		b[i] = b[j]
		i++
	}
	b = b[:i]

	if b[len(b)-1] == ' ' {
		b = b[:len(b)-1]
	}

	// reverse the whole string
	reverse(b, 0, len(b)-1)

	start := 0
	for i := range b {
		if i+1 == len(b) || b[i+1] == ' ' {
			reverse(b, start, i)
			start = i + 2
		}
	}

	return string(b)
}

func reverse(b []byte, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
