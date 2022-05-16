package _202_happy_number

// tptl passed
func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = calcSquareSum(slow)
		fast = calcSquareSum(calcSquareSum(fast))
		if slow == fast {
			break
		}
	}

	return slow == 1
}

func calcSquareSum(n int) int {
	res := 0
	for n > 0 {
		digit := n % 10
		res += digit * digit
		n /= 10
	}

	return res
}

func isHappyNaive(n int) bool {
	set := make(map[int]bool)
	for {
		n = calcSquareSum(n)
		if n == 1 {
			return true
		}
		if set[n] {
			return false
		}
		set[n] = true
	}
	return false
}
