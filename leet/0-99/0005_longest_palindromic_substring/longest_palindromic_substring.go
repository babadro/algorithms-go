package _005_longest_palindromic_substring

// TODO 2 this is brute force. Implement good solution (in description maybe)
func longestPalindrome(s string) string {
	maxPalindromeLen := 0
	var left, right int
	for i := range s {
		for j := len(s); j > i; j-- {
			if currLen := j - i; currLen > maxPalindromeLen {
				if isPalindrome(s, i, j) {
					maxPalindromeLen = currLen
					left, right = i, j
				}
			} else {
				break
			}
		}
	}

	return s[left:right]
}

func isPalindrome(s string, left, right int) bool {
	for i, j := left, right-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func longestPalindrome3(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	start := 0
	end := 0
	for j := 1; j < len(s); j++ {
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > end-start+1 {
					start = i
					end = j
				}
			}
		}
	}

	return s[start : end+1]
}
