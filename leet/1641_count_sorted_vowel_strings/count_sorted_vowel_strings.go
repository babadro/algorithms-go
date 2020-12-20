package _1641_count_sorted_vowel_strings

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}
var idxes = []int{0: 'a', 1: 'e', 2: 'i', 3: 'o', 4: 'u'}

// passed
// todo 1: understand why idxes doesn't work.
func countVowelStrings(n int) int {
	return count(n, 'a')
}

func count(n int, lastChar byte) (res int) {
	if n == 0 {
		return 1
	}
	for idx := 0; idx < len(vowels); idx++ {
		if lastChar > vowels[idx] {
			continue
		}
		res += count(n-1, vowels[idx])
	}

	return res
}
