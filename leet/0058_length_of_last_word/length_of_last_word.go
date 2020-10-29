package _058_length_of_last_word

// tptl
// 100% 100%
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	last, curr := 0, 0
	for i := range s {
		if s[i] == ' ' {
			if curr != 0 {
				last = curr
				curr = 0
			}
		} else {
			curr++
		}
	}

	if s[len(s)-1] != ' ' {
		return curr
	}

	return last
}
