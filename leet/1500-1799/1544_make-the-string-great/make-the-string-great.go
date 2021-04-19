package _1544_make_the_string_great

// passed. easy
func makeGood(s string) string {
	var res []byte

	for i := range s {
		char := s[i]
		res = append(res, char)
		n := len(res)
		if n > 1 {
			l, r := res[n-2], res[n-1]
			if l-32 == r || r-32 == l {
				res = res[:n-2]
			}
		}
	}

	return string(res)
}
