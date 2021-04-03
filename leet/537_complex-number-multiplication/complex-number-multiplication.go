package _537_complex_number_multiplication

import "strconv"

// passed. tptl
func complexNumberMultiply(a string, b string) string {
	r1, i1 := parseComplex(a)
	r2, i2 := parseComplex(b)

	r := r1*r2 - i1*i2
	i := i1*r2 + r1*i2

	return complexToString(r, i)
}

func complexToString(r, i int) string {
	res := make([]byte, 0, 9)

	res = strconv.AppendInt(res, int64(r), 10)
	res = append(res, '+')
	res = strconv.AppendInt(res, int64(i), 10)
	res = append(res, 'i')

	return string(res)
}

func parseComplex(num string) (real int, imaginary int) {
	i := 0
	if num[i] == '-' || num[i] == '+' {
		i = 1
	}
	for ; num[i] >= '0' && num[i] <= '9'; i++ {
	}
	real, _ = strconv.Atoi(num[:i])

	i++
	startImag := i
	if num[i] == '-' || num[i] == '+' {
		i++
	}

	for ; num[i] >= '0' && num[i] <= '9'; i++ {
	}

	imaginary, _ = strconv.Atoi(num[startImag:i])

	return real, imaginary
}
