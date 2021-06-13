package _392_is_subsequence

// tptl. passed. best solution. #string
func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(t) && j < len(s); i++ {
		if t[i] == s[j] {
			j++
		}
	}

	return j == len(s)
}

func isSubsequenceWithRunes(s string, t string) bool {
	i, j := 0, 0
	runesS, runesT := []rune(s), []rune(t)
	for ; i < len(runesS); i++ {
		for ; j < len(runesT) && runesS[i] != runesT[j]; j++ {
		}
		if j == len(runesT) {
			break
		}
		j++
	}

	return i == len(runesS)
}

func isSubsequenceWithoutRunes(s string, t string) bool {
	i, j := 0, 0
	for ; i < len(s); i++ {
		for ; j < len(t) && s[i] != t[j]; j++ {
		}
		if j == len(t) {
			break
		}
		j++
	}

	return i == len(s)
}
