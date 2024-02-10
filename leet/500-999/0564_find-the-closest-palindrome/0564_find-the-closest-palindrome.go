package _564_find_the_closest_palindrome

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	b := []byte(n)

	if len(b)%2 == 0 {
		return strconv.Itoa(makePalindrome(b))
	}

	candidate1 := makePalindrome(b)

	candidate2, candidate3 := math.MinInt64, math.MinInt64

	mid := len(b) / 2

	// increment middle
	if b[mid] == '9' {
		b[mid] = '0'

		for i := mid - 1; i >= 0; i-- {
			if b[i] < '9' {
				b[i]++
				break
			}

			b[i] = '0'
		}
	} else {
		b[mid]++
	}

	if b[0] == '0' {
		candidate2 = makePalindrome(b)
	}

	// decrement middle
	b = []byte(n)

	if b[mid] == '0' {
		b[mid] = '9'

		for i := mid - 1; i >= 0; i-- {
			if b[i] > '0' {
				b[i]--
				break
			}

			b[i] = '9'
		}
	}

}

func makePalindrome(b []byte) int {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[j] = b[i]
	}

	num, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}

	return num
}
