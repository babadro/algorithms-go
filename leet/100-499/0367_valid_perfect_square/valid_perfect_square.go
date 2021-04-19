package _367_valid_perfect_square

// passed. tptl. best solution
func isPerfectSquare2(num int) bool {
	l, r := 1, num
	for l <= r {
		mid := (r-l)/2 + l
		if res := mid * mid; res < num {
			l = mid + 1
		} else if res > num {
			r = mid - 1
		} else {
			return true
		}
	}

	return false
}

// bruteforce
func isPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		if res := i * i; res == num {
			return true
		} else if res > num {
			break
		}
	}

	return false
}
