package _1417_reformat_the_string

// todo 1
func reformat(s string) string {
	n := len(s)

	digitsCount := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			digitsCount++
		}
	}

	if n%2 == 0 {
		if digitsCount != n/2 {
			return ""
		}
	} else if !(digitsCount == n/2 || digitsCount == n/2+1) {
		return ""
	}

}
