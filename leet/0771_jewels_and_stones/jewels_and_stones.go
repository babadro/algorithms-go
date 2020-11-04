package _771_jewels_and_stones

// easy. passed.
func numJewelsInStones(J string, S string) int {
	m := make(map[int32]bool)
	for _, char := range J {
		m[char] = true
	}

	counter := 0
	for _, char := range S {
		if m[char] {
			counter++
		}
	}

	return counter
}
