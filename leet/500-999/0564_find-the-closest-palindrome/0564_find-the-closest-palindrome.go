package _564_find_the_closest_palindrome

import (
	"math"
	"strconv"
)

// #bnsrg hard passed
// todo 2 find shorter solution
func nearestPalindromic(n string) string {
	if len(n) <= 2 {
		return bruteForce(n)
	}

	b := []byte(n)

	candidate1, candidate2, candidate3 := math.MinInt64, math.MinInt64, math.MinInt64

	if !isPalindrome(b) {
		candidate1 = makePalindromeFromLeftSide(b)
	}

	leftLen, rightLen := len(n)/2, len(n)/2
	if len(n)%2 == 1 {
		leftLen += 1
	}

	leftPart, err := strconv.Atoi(n[:leftLen])
	if err != nil {
		panic(err)
	}

	incremented, decremented := strconv.Itoa(leftPart+1), strconv.Itoa(leftPart-1)

	candidate2 = makePalindromeFromLeftSide(append([]byte(incremented), make([]byte, rightLen)...))

	candidate3 = makePalindromeFromLeftSide(append([]byte(decremented), make([]byte, rightLen)...))

	var res int

	num, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	diff1, diff2, diff3 := diff(num, candidate1), diff(num, candidate2), diff(num, candidate3)

	switch {
	case diff1 < diff2:
		res = candidate1
	case diff1 == diff2:
		res = min(candidate1, candidate2)
	default:
		res = candidate2
	}

	diffRes := diff(num, res)
	if diffRes == diff3 {
		res = min(res, candidate3)
	} else if diff3 < diffRes {
		res = candidate3
	}

	return strconv.Itoa(res)
}

func diff(a, b int) int {
	if d := a - b; d < 0 {
		return -d
	} else {
		return d
	}
}

func makePalindromeFromLeftSide(b []byte) int {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[j] = b[i]
	}

	if mid := len(b) / 2; b[mid] == 0 {
		b[mid] = b[0]
	}

	num, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}

	return num
}

func isPalindrome(b []byte) bool {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return false
		}
	}

	return true
}

var palindromes = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 22, 33, 44, 55, 66, 77, 88, 99, 101}

func bruteForce(n string) string {
	num, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	minDiff, res := math.MaxInt64, 0
	for _, palindrome := range palindromes {
		d := diff(num, palindrome)
		if d == 0 {
			continue
		} else if d < minDiff {
			minDiff, res = d, palindrome
		} else if d == minDiff {
			res = min(res, palindrome)
		}
	}

	return strconv.Itoa(res)
}
