package _044_wildcard_matching

// todo 1
func isMatch(s string, p string) bool {

}

func match(b []byte, pattern byte) bool {
	if isLetter(pattern) {
		return len(b) == 1 && b[0] == pattern
	}

	if pattern == '?' {
		return len(b) == 1
	}

	return true
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z'
}
