package _844_backspace_string_compare

func backspaceCompare3(S string, T string) bool {
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
