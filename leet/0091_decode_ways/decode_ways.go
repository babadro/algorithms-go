package _091_decode_ways

// TODO 1
func numDecodings(s string) int {
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
	if curr == '1' || (curr == '2' && next >= '0' && next <= '6') {
		*counter++
		decode(s, n, idx+2, counter)
	}
}
