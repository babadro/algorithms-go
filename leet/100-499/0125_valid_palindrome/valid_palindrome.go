package _125_valid_palindrome

// tptl. passed.
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if i < j && !isAlphanumeric(s[i]) {
			i++
			continue
		}

		if j > i && !isAlphanumeric(s[j]) {
			j--
			continue
		}

		if toLowercase(s[i]) != toLowercase(s[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(char byte) bool {
	return (char >= '0' && char <= '9') || (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}

func toLowercase(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + 32
	}

	return char
}
