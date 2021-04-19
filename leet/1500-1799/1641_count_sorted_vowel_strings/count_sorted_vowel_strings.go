package _1641_count_sorted_vowel_strings

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}
var idxes = []int{'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4}

// passed
func countVowelStrings(n int) int {
	return count(n, 'a')
}

func count(n int, lastChar byte) (res int) {
	if n == 0 {
		return 1
	}
	for idx := idxes[lastChar]; idx < len(vowels); idx++ {
		if lastChar > vowels[idx] {
			continue
		}
		res += count(n-1, vowels[idx])
	}

	return res
}
