package _1_palindrome_permutation

// leetcode premium https://leetcode.com/problems/palindrome-permutation/
func permutePalindrome(st string) bool {
	m := make(map[int32]int)

	oddCount := 0
	for _, ch := range st {
		m[ch]++
		if m[ch]%2 == 1 {
			oddCount++
		} else {
			oddCount--
		}
	}

	if len(st)%2 == 1 {
		return oddCount == 1
	}

	return oddCount == 0
}
