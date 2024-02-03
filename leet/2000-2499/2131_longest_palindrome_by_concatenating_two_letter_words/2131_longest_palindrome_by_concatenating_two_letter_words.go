package _2131_longest_palindrome_by_concatenating_two_letter_words

// #bnsrg medium passed
func longestPalindrome(words []string) int {
	m := make(map[string]int)

	for _, w := range words {
		m[w]++
	}

	res := 0
	oddExists := false

	for word, count := range m {
		if word[0] == word[1] {
			if count%2 != 0 {
				oddExists = true
				count -= 1
			}

			res += count * 2

			continue
		}

		res += min(m[word], m[string([]byte{word[1], word[0]})]) * 2
	}

	if oddExists {
		res += 2
	}

	return res
}
