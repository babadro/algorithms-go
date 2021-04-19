package _759_count_number_of_homogenous_substrings

// passed. todo 2 make shorter solution
func minOperations(s string) int {
	count1 := 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] != '1' {
				count1++
			}
		} else {
			if s[i] != '0' {
				count1++
			}
		}
	}

	count2 := 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] != '0' {
				count1++
			}
		} else {
			if s[i] != '1' {
				count1++
			}
		}
	}

	if count1 < count2 {
		return count1
	}

	return count2
}
