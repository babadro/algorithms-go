package _720_longest_word_in_dictionary

import "sort"

// brute force. passed
func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}

		return len(words[i]) > len(words[j])
	})

	wordsSet := make(map[string]bool, len(words))
	for _, w := range words {
		wordsSet[w] = true
	}

Loop:
	for _, w := range words {
		for j := 1; j <= len(w); j++ {
			if !wordsSet[w[:j]] {
				continue Loop
			}
		}

		return w
	}

	return ""
}
