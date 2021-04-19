package _844_backspace_string_compare

// tptl. best solution. passed. O(N) and O(1)
func backspaceCompare2(S string, T string) bool {
	i, j := len(S)-1, len(T)-1
	skipS, skipT := 0, 0

	for i >= 0 || j >= 0 {
		for i >= 0 {
			if S[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if T[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 && S[i] != T[j] {
			return false
		}
		if (i >= 0) != (j >= 0) {
			return false
		}
		i--
		j--
	}

	return true
}

// brutforce. passed.
func backspaceCompare(S string, T string) bool {
	var resS []rune
	var resT []rune

	for _, char := range []rune(S) {
		if char == '#' {
			if len(resS) > 0 {
				resS = resS[:len(resS)-1]
			}
		} else {
			resS = append(resS, char)
		}
	}

	for _, char := range []rune(T) {
		if char == '#' {
			if len(resT) > 0 {
				resT = resT[:len(resT)-1]
			}
		} else {
			resT = append(resT, char)
		}
	}

	return string(resS) == string(resT)
}
