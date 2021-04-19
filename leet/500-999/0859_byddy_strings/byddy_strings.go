package _859_byddy_strings

// passed. tptl
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) || len(A) < 2 {
		return false
	}

	rA := []rune(A)

	if A == B {
		m := make(map[int32]bool, len(rA))
		for i := range rA {
			if !m[rA[i]] {
				m[rA[i]] = true
			} else {
				return true
			}
		}

		return false
	}

	rB := []rune(B)
	chars := make([]int32, 0, 4)
	counter := 0

	for i := range rA {
		if rA[i] != rB[i] {
			if counter > 2 {
				return false
			}
			chars = append(chars, rA[i], rB[i])
			counter++
		}
	}

	if counter != 2 {
		return false
	}

	if chars[0] != chars[3] || chars[1] != chars[2] {
		return false
	}

	return true
}
