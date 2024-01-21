package _1062_longest_repeating_substring

// #medium passed, but slow
// todo 2 fast dp solution exists
func longestRepeatingSubstring(s string) int {
	m := make(map[string]bool)

	for n := len(s) - 1; n >= 2; n-- {
		for i := 0; i+n <= len(s); i++ {
			subStr := s[i : i+n]

			if m[subStr] {
				return n
			}

			m[subStr] = true
		}
	}

	return 0
}
