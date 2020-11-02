package _221_split_string_in_balanced_strings

// tptl. easy
func balancedStringSplit(s string) int {
	balance, amount := 0, 0

	for _, char := range s {
		if char == 'L' {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			amount++
		}
	}

	return amount
}
