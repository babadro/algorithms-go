package _1930_unique_length_3_palindromic_subsequences

// tptl passed #string #medium best solution
func countPalindromicSubsequence(s string) int {
	m1, m2 := [26]int{}, [26]int{}

	for i := 0; i < len(s); i++ {
		m2[s[i]-'a'] = i
	}
	for i := len(s) - 1; i >= 0; i-- {
		m1[s[i]-'a'] = i
	}

	res := make(map[[2]byte]bool)

	for i := 0; i < 26; i++ {
		char1 := byte('a' + i)
		if m1[i] != m2[i] && (m2[i]-m1[i] >= 2) {
			for j := m1[i] + 1; j < m2[i]; j++ {
				char2 := s[j]

				res[[2]byte{char1, char2}] = true
			}
		}
	}

	return len(res)
}
