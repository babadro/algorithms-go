package _1930_unique_length_3_palindromic_subsequences

// todo 1. TLE
func countPalindromicSubsequenceBruteForce(s string) int {
	uniq := [26]bool{}
	count := 0
	for i := range s {
		idx := s[i] - 'a'
		if !uniq[idx] {
			uniq[idx] = true
			count++

			if count == 676 {
				break
			}
		}
	}

	max := count * count

	m := make(map[[2]byte]bool)
	tmp := [2]byte{}

	for i := 0; i < len(s)-2; i++ {
		tmp[0] = s[i]
		for j := i + 1; j < len(s)-1; j++ {
			tmp[1] = s[j]
			if m[tmp] {
				continue
			}

			for k := j + 1; k < len(s); k++ {
				if s[k] == tmp[0] {
					m[tmp] = true

					if len(m) == max {
						return max
					}

					break
				}
			}
		}
	}

	return len(m)
}
