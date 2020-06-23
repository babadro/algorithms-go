package _091_decode_ways

// TODO 1
func numDecodings(s string) int {
	return decode(s, len(s), 0, 1)
}

func decode(s string, n, idx, counter int) int {
	if idx >= n-1 {
		return counter
	}
	counter += decode(s, n, idx+1, counter)
	curr, next := s[idx], s[idx+1]
	if curr == '1' || (curr == '2' && next >= '0' && next <= '6') {
		counter += decode(s, n, idx+2, counter+1)
	}
	return counter
}
