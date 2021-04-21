package _679_24_game

import "strconv"

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
	n, nb := len(nums), len(brackets)

	if  n == 0 && nb == 0 {
		return eval(stmt) == 24
	}

	// nums
	if n == 4 {
		for i := 0; i < 4; i++ {
			nextStmt := append(stmt, strconv.Itoa(nums[i])...)
			nums[i], nums[n-1] = nums[n-1], nums[i]
			nextNums := nums[:n-1]

			if helper(nextNums, nextStmt, brackets) {
				return true
			}
		}
	} else if ...{}


	// brackets
	if n == 4 {
		if nb > 0 && brackets[0] == '(' {
			nextStmt := append(stmt, '(')
			nextBrackets := brackets[1:]

			return helper(nums, nextStmt, nextBrackets)
		}
	}




	return false
}

func eval(stmt []byte) int {
	return 0
}
