package _679_24_game

import "fmt"

var operators = []byte{'*', '/', '+', '-'}
var brackets = []string{"", "()", "()()", "(())"}

// todo 1
func judgePoint24(nums []int) bool {
	numPermutations := genPerm(nums)

	stmt := make([]byte, 0, 11)
	for _, br := range brackets {
		for _, perm := range numPermutations {
			if helper(perm, stmt, []byte(br)) {
				return true
			}
		}
	}

	return false
}

func helper(nums string, stmt []byte, brackets []byte) bool {
	n, nb, ns := len(nums), len(brackets), len(stmt)

	if n == 0 && nb == 0 {
		return eval(stmt) == 24
	}

	last := ns - 1
	// nums
	if n > 0 && (n == 4 || stmt[last] == '(' || isOperator(stmt[last])) &&
		!(n == 2 && nb == 2) && !(n == 3 && nb == 4) && !(n == 2 && nb == 3) {
		nextStmt := append(stmt, nums[0])
		nextNums := nums[1:]

		if helper(nextNums, nextStmt, brackets) {
			return true
		}
	}

	// brackets
	if !(nb == 0 ||
		(ns > 1 && stmt[ns-2] == '(' && brackets[0] == ')') ||
		(ns > 0 && isOperator(stmt[last])) ||
		(ns > 0 && isNum(stmt[last]) && brackets[0] == '(') ||
		(ns > 0 && stmt[last] == ')' && brackets[0] == '(') ||
		(ns > 0 && stmt[last] == '(' && brackets[0] == ')')) {
		nextStmt := append(stmt, brackets[0])
		nextBrackets := brackets[1:]

		if helper(nums, nextStmt, nextBrackets) {
			return true
		}
	}

	// operators
	if !(n == 0 || n == 4 || (ns > 0 && (isOperator(stmt[last]) || stmt[last] == '('))) {
		for _, operator := range operators {
			nextStmt := append(stmt, operator)

			if helper(nums, nextStmt, brackets) {
				return true
			}
		}
	}

	return false
}

func isOperator(b byte) bool {
	for _, op := range operators {
		if b == op {
			return true
		}
	}

	return false
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func eval(stmt []byte) (float64, bool) {
	n := len(stmt)
	for n > 1 && stmt[0] == '(' && stmt[n-1] == ')' {
		stmt = stmt[1 : n-1]
		n = len(stmt)
	}

	res, operator := float64(0), byte('+')
	for i, j := 0, 1; i <= len(stmt); i = j + 1 {
		bracketBalance := 0
		for ; j < n && (bracketBalance > 0 || (stmt[j] != '+' && stmt[j] != '-')); j++ {
			if stmt[j] == '(' {
				bracketBalance++
			} else if stmt[j] == ')' {
				bracketBalance--
			}
		}

		next, ok := eval(stmt[i:j])
		if !ok {
			return 0, false
		}

		if res, ok = exec(res, next, operator); !ok {
			return 0, false
		}

		if j < n {
			operator = stmt[j]
		}
	}

	return 0, true
}

func exec(l, r float64, op byte) (float64, bool) {
	switch op {
	case '+':
		return l + r, true
	case '-':
		return l - r, true
	case '/':
		if r == 0 {
			return 0, false
		}

		return l / r, true
	case '*':
		return l * r, true
	}

	panic(fmt.Sprintf("unknown operator %s", string(op)))
}

func genPerm(arr []int) []string {
	var res []string
	rec(arr, &res, 0, len(arr)-1)

	return res
}

func rec(arr []int, res *[]string, l, r int) {
	if l == r {
		b := make([]byte, len(arr))
		for i, num := range arr {
			b[i] = byte('0' + num)
		}

		*res = append(*res, string(b))
	}

	for i := l; i <= r; i++ {
		arr[i], arr[l] = arr[l], arr[i]
		rec(arr, res, l+1, r)
		arr[i], arr[l] = arr[l], arr[i]
	}
}
