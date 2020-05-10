package _005_longest_palindromic_substring

// TODO 2 this is brute force. Implement good solution (in description maybe)
func longestPalindrome(s string) string {
	runes := []rune(s)
	maxPalindromeLen := 0
	var current []rune
	for i := range runes {
		for j := len(runes); j > i; j-- {
			if currLen := j - i; currLen > maxPalindromeLen {
				candidate := runes[i:j]
				if isPalindrome(candidate) {
					maxPalindromeLen = currLen
					current = candidate
				}
			}
		}
	}
	return string(current)
}

func isPalindrome(runes []rune) bool {
	l := len(runes)
	for i := 0; i < l/2; i++ {
		if runes[i] != runes[l-1-i] {
			return false
		}
	}
	return true
}
