package _022_generate_parentheses

import "math"

// Runtime: 4 ms, faster than 16.75% of Go online submissions for Generate Parentheses.
// Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Generate Parentheses.
func generateParenthesis1(n int) []string {
	var res []string
	for i := 0; i <= int(math.Pow(2, 2*float64(n))-1); i++ {
		if !valid1(i, 2*n) {
			continue
		}
		res = append(res, makeParenthesis(i, 2*n))
	}
	return res
}

func valid1(num, parenthesisCount int) bool {
	stack := make([]bool, 0)
	for i := 0; i < parenthesisCount; i++ {
		if num&1 == 1 {
			stack = append(stack, true)
		} else if l := len(stack); l == 0 {
			return false
		} else {
			stack = stack[:l-1]
		}
		num = num >> 1
	}
	return len(stack) == 0
}

func makeParenthesis(num, parenthesisCount int) string {
	arr := make([]byte, 0, parenthesisCount)
	for i := 0; i < parenthesisCount; i++ {
		if num&1 == 1 {
			arr = append(arr, '(')
		} else {
			arr = append(arr, ')')
		}
		num = num >> 1
	}
	return string(arr)
}

// redundant
func getMax(n int) int {
	max := 1<<(n*2) - 1
	for i := 0; i < n; i++ {
		max &^= 1 << i
	}
	return max
}

func bitCount(n int) int {
	res := 0
	for n != 0 {
		if n&1 == 1 {
			res++
		}
		n = n >> 1
	}
	return res
}
