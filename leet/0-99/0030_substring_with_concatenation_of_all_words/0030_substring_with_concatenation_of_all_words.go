package _030_substring_with_concatenation_of_all_words

// tptl. passed, but not the best.
// todo 2 check fastest solution
func findSubstring(s string, words []string) []int {
	freq := make(map[string]int)

	for _, w := range words {
		freq[w]++
	}

	wLen := len(words[0])
	var res []int
	for i := 0; i <= len(s)-len(words)*wLen; i++ {
		visited := make(map[string]int)
		for j := 0; j < len(words); j++ {
			idx := i + j*wLen
			word := s[idx : idx+wLen]
			count := freq[word]
			if count == 0 || visited[word] == count {
				break
			}

			visited[word]++
			if j == len(words)-1 {
				res = append(res, i)
			}
		}
	}

	return res
}
