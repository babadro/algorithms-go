package _1556_thousand_separator

// tptl. passed.
func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	res := make([]byte, 0, 9)
	for i := 0; n != 0; i++ {
		if i > 0 && i%3 == 0 {
			res = append(res, '.')
		}

		r := byte(n % 10)
		res = append(res, '0'+r)

		n = n / 10
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return string(res)
}
