package _1332_remove_palindromic_subsequences

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	if IsPalindrome(s) {
		return 1
	}
	return 2
}

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
