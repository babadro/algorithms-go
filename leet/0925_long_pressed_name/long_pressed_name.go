package _925_long_pressed_name

// past best solution. tptl not mine.
func isLongPressedName2(name string, typed string) bool {
	i, m, n := 0, len(name), len(typed)
	for j := 0; j < n; j++ {
		if i < m && name[i] == typed[j] {
			i++
		} else if j == 0 || typed[j] != typed[j-1] {
			return false
		}
	}
	return i == m
}

// passed
func isLongPressedName(name string, typed string) bool {
	n, m := len(name), len(typed)
	var i, j int
	for i, j = 0, 0; i < n && j < m; {
		origCounter, typedCounter := 1, 0
		for i < n-1 && name[i] == name[i+1] {
			origCounter++
			i++
		}

		for ; j < m && typed[j] == name[i]; j++ {
			typedCounter++
		}

		if origCounter > typedCounter {
			return false
		}

		i++
	}

	return i == n && j == m
}
