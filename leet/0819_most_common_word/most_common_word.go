package _819_most_common_word

import "strings"

// passed. easy
func mostCommonWord(paragraph string, banned []string) string {
	m := make(map[string]int)
	bannedDict := make(map[string]bool, len(banned))

	for _, ban := range banned {
		bannedDict[ban] = true
	}

	start := 0
	for i := 0; i < len(paragraph); {
		char := paragraph[i]

		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if i == len(paragraph)-1 {
				i++
			} else {
				continue
			}
		}

		if start < i {
			word := strings.ToLower(paragraph[start:i])
			if !bannedDict[word] {
				m[word]++
			}
		}

		start = i + 1
	}

	word, max := "", 0
	for key, value := range m {
		if value > max {
			max = value
			word = key
		}
	}

	return word
}
