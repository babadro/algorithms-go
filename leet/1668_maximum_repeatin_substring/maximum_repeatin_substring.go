package _1668_maximum_repeatin_substring

// passed. tptl
func maxRepeating(sequence string, word string) int {
	n := len(word)
	max := 0
	for k := 0; k+n <= len(sequence); k++ {
		counter := 0
		for i := k; i+n <= len(sequence) && sequence[i:i+n] == word; i += n {
			counter++
		}

		if counter > max {
			max = counter
		}
	}

	return max
}
