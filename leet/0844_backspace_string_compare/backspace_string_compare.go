package _844_backspace_string_compare

// brutforce. passed.
// todo 1 solve it with O(N) and O(1)
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
