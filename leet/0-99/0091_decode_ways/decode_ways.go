package _091_decode_ways

// 100%, 93%
func numDecodings(s string) int {
	n := len(s)
	if s[0] == '0' {
		return 0
	}

	arr := make([]int, n)
	if last := n - 1; s[last] != '0' {
		arr[last] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if nextNum := arr[i+1]; nextNum == 0 {
			arr[i] = 1
		} else {
			arr[i] = nextNum
		}

		currChar, nextChar := s[i], s[i+1]
		if nextChar == '0' {
			if currChar == '0' || currChar > '2' {
				return 0
			}

			continue
		}

		idx := i + 2
		if (currChar == '1' && nextChar > '0') || (currChar == '2' && nextChar > '0' && nextChar <= '6') {
			if idx >= n {
				arr[i] = 2
			} else {
				arr[i] += arr[idx]
			}
		}
	}

	return arr[0]
}

func numDecodingsRecursive(s string) int {
	counter := 1
	decode(s, len(s), 0, &counter)
	return counter
}

func decode(s string, n, idx int, counter *int) {
	if idx >= n-1 {
		return
	}
	decode(s, n, idx+1, counter)
	curr, next := s[idx], s[idx+1]
	if curr == '1' || (curr == '2' && next >= '0' && next <= '6') { // incorrect conditional.
		*counter++
		decode(s, n, idx+2, counter)
	}
}
