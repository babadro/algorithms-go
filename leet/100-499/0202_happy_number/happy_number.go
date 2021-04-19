package _202_happy_number

import "fmt"

func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = calcSquareSum(slow)
		fast = calcSquareSum(calcSquareSum(fast))
		fmt.Println(slow, " ", fast)
		if slow == fast {
			break
		}
	}
	return slow == 1
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

func calcSquareSum(n int) int {
	res, digit := 0, 0
	for n > 0 {
		digit = n % 10
		res += digit * digit
		n = n / 10
	}
	return res
}
