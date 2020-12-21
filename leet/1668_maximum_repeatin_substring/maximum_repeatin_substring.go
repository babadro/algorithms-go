package _1668_maximum_repeatin_substring

// todo 1
func maxRepeating(sequence string, word string) int {
	n := len(word)
	max := 0
	for i := 0; i+n <= len(sequence); {
		if sequence[i:i+n] == word {
			counter, j := 1, i+n
			for ; j+n <= len(sequence) && sequence[j:j+n] == word; j += n {
				counter++
			}

			if counter > max {
				max = counter
			}

			i = j + 1
		} else {
			i++
		}
	}

	return max
}
