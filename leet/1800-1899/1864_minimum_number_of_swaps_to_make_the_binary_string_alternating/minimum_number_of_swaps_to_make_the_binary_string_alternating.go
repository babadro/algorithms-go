package _1864_minimum_number_of_swaps_to_make_the_binary_string_alternating

// best solution, but hard to understand
func minSwaps(s string) int {
	zeroes, ones := 0, 0
	for i := range s {
		if s[i] == '0' {
			zeroes++
		}
	}

	ones = len(s) - zeroes
	if diff := ones - zeroes; diff > 1 || diff < -1 {
		return -1
	}

	miss0, miss1 := 0, 0
	for i := 0; i < len(s); i += 2 {
		if s[i] != '0' {
			miss0++
		} else {
			miss1++
		}
	}

	if zeroes == ones {
		return min(miss0, miss1)
	}

	if zeroes > ones {
		return miss0
	}

	return miss1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// passed. tptl #medium #string
func minSwaps2(s string) int {
	zeroes, ones := 0, 0
	for i := range s {
		if s[i] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	odd := len(s)%2 == 1
	if odd {
		if diff := zeroes - ones; diff != 1 && diff != -1 {
			return -1
		}

		startWith := '0'
		if ones > zeroes {
			startWith = '1'
		}

		return swapCounter(s, byte(startWith))
	} else if zeroes != ones {
		return -1
	}

	count1, count2 := swapCounter(s, '0'), swapCounter(s, '1')
	if count1 < count2 {
		return count1
	}

	return count2
}

func swapCounter(s string, startWith byte) int {
	res := 0
	for i := range s {
		if i%2 == 0 {
			if s[i] != startWith {
				res++
			}
		} else if s[i] == startWith {
			res++
		}
	}

	return res / 2
}

// doesn't work
func minSwapsDoesntWork(s string) int {
	count1 := swapCount([]byte(s), '0')
	count2 := swapCount([]byte(s), '1')

	if count1 < count2 {
		return count1
	}

	return count2
}

func swapCount(b []byte, startWith byte) int {
	n := len(b)
	ans := 0
	for i, k := 0, 0; i < n; i++ {
		if i%2 == 1 && b[i] == startWith {
			for ; k < n; k++ {
				if b[k] == startWith || k%2 == 1 {
					continue
				}

				break
			}

			if k == n {
				return -1
			}

			b[k], b[i] = b[i], b[k]
			ans++
		}
	}

	return ans
}
