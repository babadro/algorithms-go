package _044_wildcard_matching // todo 1
import "fmt"

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	if len(p) == 0 {
		fmt.Printf("%s, %s\n", s, p)
		return false
	}

	char, nextP := p[0], p[1:]

	if char >= 'a' && char <= 'z' {
		if len(s) == 0 || s[0] != char {
			fmt.Printf("%s, %s, %s\n", s, p, string(char))
			return false
		}

		return isMatch(s[1:], nextP)
	}

	if char == '?' {
		if len(s) == 0 {
			fmt.Printf("%s, %s, %s\n", s, p, string(char))
			return false
		}

		return isMatch(s[1:], nextP)
	}

	// char == *
	if len(nextP) == 0 {
		return true
	}

	for i := 0; i <= len(s); i++ {
		if isMatch(s[i:], nextP) {
			return true
		}
	}

	fmt.Printf("%s, %s, %s\n", s, p, string(char))
	return false
}
