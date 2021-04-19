package _792_number_of_matching_subsequences

// bruteforce. passed, but slowly.
func numMatchingSubseq(S string, words []string) int {
	counter := 0
	for _, v := range words {
		if isSubsequence(S, v) {
			counter++
		}
	}

	return counter
}

func isSubsequence(s, t string) bool {
	i, j := 0, 0
	for ; i < len(s) && j < len(t); i++ {
		if s[i] == t[j] {
			j++
		}
	}

	return j == len(t)
}
