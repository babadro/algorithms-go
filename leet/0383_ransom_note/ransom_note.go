package _0383_ransom_note

// passed. easy. tptl
func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 256)
	for _, letter := range magazine {
		m[letter]++
	}

	for _, letter := range ransomNote {
		if m[letter] == 0 {
			return false
		}

		m[letter]--
	}

	return true
}
