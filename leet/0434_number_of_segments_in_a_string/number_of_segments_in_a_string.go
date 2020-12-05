package _434_number_of_segments_in_a_string

// passed. tptl. best solution (shorter, than first solution)
func countSegments2(s string) int {
	counter := 0
	for i := 0; i < len(s); i++ {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			counter++
		}
	}

	return counter
}

// passed, but not shortest.
func countSegments(s string) int {
	counter, flag := 0, false
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if !flag {
				flag = true
				counter++
			}
		} else {
			flag = false
		}
	}

	return counter
}
