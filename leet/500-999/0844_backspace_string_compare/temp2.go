package _844_backspace_string_compare

func backspaceCompare4(S string, T string) bool {
	i, j := len(S)-1, len(T)-1
	skipS, skipT := 0, 0
	for i >= 0 || j >= 0 {
		if i >= 0 {
			if S[i] == '#' {
				skipS++
				i--
				continue
			} else if skipS > 0 {
				skipS--
				i--
				continue
			}
		}

		if j >= 0 {
			if T[j] == '#' {
				skipT++
				j--
				continue
			} else if skipT > 0 {
				skipT--
				j--
				continue
			}
		}

		if i < 0 || j < 0 || S[i] != T[j] {
			return false
		}

		i--
		j--
	}

	return i < 0 && j < 0
}
