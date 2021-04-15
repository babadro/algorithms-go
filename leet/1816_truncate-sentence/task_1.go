package _816_truncate_sentence

import "strings"

// passed. easy.
func truncateSentence(s string, k int) string {
	words := strings.Split(s, " ")

	res := make([]byte, 0)

	for i := 0; i < k; i++ {
		word := words[i]

		if i > 0 {
			res = append(res, ' ')
		}

		res = append(res, word...)
	}

	return string(res)
}
