package _1324_print_words_vertically

import "strings"

func printVertically(s string) []string {
	input := strings.Split(s, " ")
	resLen := 0
	for _, v := range input {
		if length := len(v); length > resLen {
			resLen = length
		}
	}

	res := make([]string, 0, resLen)
	for j := 0; j < resLen; j++ {
		word := make([]byte, 0, len(input))
		for i := range input {
			if j < len(input[i]) {
				word = append(word, input[i][j])
			} else {
				word = append(word, " "...)
			}
		}
		trailedWord := strings.TrimRight(string(word), " ")
		res = append(res, trailedWord)
	}
	return res
}
