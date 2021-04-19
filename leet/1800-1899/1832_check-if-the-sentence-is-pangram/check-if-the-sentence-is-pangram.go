package _1832_check_if_the_sentence_is_pangram

// passed. easy
func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}

	dict := [26]bool{}
	counter := 0
	for i := range sentence {
		idx := sentence[i] - 'a'
		if !dict[idx] {
			dict[idx] = true
			counter++

			if counter == 26 {
				break
			}
		}
	}

	return counter == 26
}
