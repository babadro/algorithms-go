package _043_multiply_strings

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	l1, l2 := len(num1), len(num2)
	shortNum, longNum := num1, num2
	if l2 < l1 {
		shortNum, longNum = num2, num1
	}

	ld := make([]int, len(longNum))
	for j := 0; j < len(ld); j++ {
		ld[j] = int(longNum[j] - '0')
	}

	result := make([]int, (l1+l2)*2)
	lastDigit := 0

	for j, start := len(shortNum)-1, 0; j >= 0; j, start = j-1, start+1 {
		lastDigit = mult(ld, int(shortNum[j]-'0'), result, start)
	}

	var sb strings.Builder

	// reverse
	for j := lastDigit; j >= 0; j-- {
		sb.WriteString(strconv.Itoa(result[j]))
	}

	return sb.String()
}

func mult(a []int, n int, result []int, start int) int {
	carryOver := 0

	for j := len(a) - 1; j >= 0; j, start = j-1, start+1 {
		res := (a[j] * n) + carryOver + result[start]
		if res > 9 {
			carryOver = res / 10
			result[start] = res - (carryOver * 10)
		} else {
			carryOver = 0
			result[start] = res
		}
	}

	for ; carryOver > 0; start++ {
		res := carryOver + result[start]
		if res > 9 {
			carryOver = res / 10
			result[start] = res - (carryOver * 10)
		} else {
			carryOver = 0
			result[start] = res
		}
	}

	return start - 1
}
