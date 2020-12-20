package _1641_count_sorted_vowel_strings

func countVowelStrings2(n int) int {
	return count2(n, 0)
}

func count2(n, last int) int {
	if n == 0 {
		return 1
	}

	var res int
	for i := last; i < 5; i++ {
		res += count2(n-1, i)
	}

	return res
}
