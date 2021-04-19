package _412_fizz_buzz

import "strconv"

const (
	fizz  = "Fizz"
	buzz  = "Buzz"
	fBuzz = "FizzBuzz"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)
	var item string
	var f, b bool
	for i := 1; i <= n; i++ {
		f = i%3 == 0
		b = i%5 == 0
		if f && b {
			item = fBuzz
		} else if f {
			item = fizz
		} else if b {
			item = buzz
		} else {
			item = strconv.Itoa(i)
		}
		res[i-1] = item
	}
	return res
}
