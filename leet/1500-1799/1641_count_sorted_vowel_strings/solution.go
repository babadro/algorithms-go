package _1641_count_sorted_vowel_strings

// passed. recursive
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

// passed. best solution. tptl
func countVowelStrings3(n int) int {
	// arr[i] means count of count of words ends at a, e, i, o, u
	arr := make([]int, 5)
	for i := range arr {
		arr[i] = 1
	}

	for i := 0; i < n-1; i++ {
		next := make([]int, 5)

		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				next[j] += arr[k]
			}

		}

		arr = next
	}

	var res int
	for i := range arr {
		res += arr[i]
	}
	return res
}
