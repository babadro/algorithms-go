package _139_word_break

import "strings"

// draft
func wordBreakDraft(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		if strings.Index(s, word) != -1 {
			dict[word] = true
		}
	}
	if len(dict) == 0 {
		return false
	}

	// some code
	return false

}
