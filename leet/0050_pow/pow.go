package _050_pow

// Brutforce
// Time Limit Exceeded 1.00000 -2147483648
func myPowNaive(x float64, n int) float64 {
	res := x
	if n == 0 {
		return 1
	}
	negative := n < 0
	if negative {
		n = -n
	}
	for n > 1 {
		res *= x
		n--
	}
	if negative {
		return 1 / res
	}
	return res
}

// Best solution
// TODO 2 need to understand
func myPow(x float64, n int) float64 {
	if n < 0 {
		x, n = 1/x, -n
	}
	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x, n = x*x, n>>1
	}
	return result
}

func myPowRecursive(x float64, n int) float64 {

	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / myPowRecursive(x, -n)
	}

	r := myPowRecursive(x, n/2)
	if n%2 == 0 {
		return r * r
	} else {
		return r * r * x
	}
}
