package _890_find_and_replace_pattern

import "github.com/babadro/algorithms-go/slices"

// passed. tptl medium
func findAndReplacePattern(words []string, pattern string) []string {
	patternArr := idxArr(pattern, make([]int, len(pattern)))

	var res []string
	var arr []int

	for _, word := range words {
		arr = idxArr(word, arr)
		if slices.IntSlicesAreEqual(patternArr, arr) {
			res = append(res, word)
		}
	}

	return res
}

func idxArr(word string, arr []int) []int {
	dict := [26]int{}
	for i := range dict {
		dict[i] = -1
	}

	arr = arr[:0]
	for i := range word {
		pos := word[i] - 'a'
		if idx := dict[pos]; idx == -1 {
			dict[pos] = i
		}

		arr = append(arr, dict[pos])
	}

	return arr
}
