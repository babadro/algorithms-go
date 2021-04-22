package _679_24_game

import (
	"fmt"
	"strconv"
)

var operators = []byte{'*', '/', '+', '-'}
var brackets = []string{"", "()", "()()", "(())"}

// todo 1
func judgePoint24(nums []int) bool {
	stmt := make([]byte, 0, 11)
	for _, br := range brackets {
		if helper(nums, stmt, []byte(br)) {
			return true
		}
	}

	return false
}

func helper(nums []int, stmt []byte, brackets []byte) bool {
	n, nb, ns := len(nums), len(brackets), len(stmt)

	if n == 0 && nb == 0 {
		return eval(stmt) == 24
	}

	last := ns - 1
	// nums
	if (n == 4 || stmt[last] == '(' || isOperator(stmt[last])) &&
		!(n == 2 && nb == 2) && !(n == 3 && nb == 4) && !(n == 2 && nb == 3) {
		for i := 0; i < n; i++ {
			nextStmt := append(stmt, strconv.Itoa(nums[i])...)
			nums[i], nums[n-1] = nums[n-1], nums[i] // todo i'm not sure
			nextNums := nums[:n-1]

			if helper(nextNums, nextStmt, brackets) {
				return true
			}
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

// todo take into account division by zero
func eval(stmt []byte) int {
	fmt.Println(string(stmt))

	return 0
}
