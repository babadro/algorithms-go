package _451_sort_characters_by_frequency

import "sort"

// passed. easy
func frequencySort(s string) string {
	b := []byte(s)
	freq := make([]int, 256)
	for _, ch := range b {
		freq[ch]++
	}

	sort.Slice(b, func(i, j int) bool {
		freqI, freqJ := freq[b[i]], freq[b[j]]
		if freqI != freqJ {
			return freqI > freqJ
		}

		return b[i] > b[j]
	})

	return string(b)
}
