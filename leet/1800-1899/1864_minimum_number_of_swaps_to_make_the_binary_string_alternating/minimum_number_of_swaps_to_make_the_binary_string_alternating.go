package _1864_minimum_number_of_swaps_to_make_the_binary_string_alternating

// todo 1 doesn't work
func minSwaps(s string) int {
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
		if i%2 == 0 && b[i] == startWith {
			for ; k < n; k++ {
				if b[k] == startWith || k%2 == 0 {
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
