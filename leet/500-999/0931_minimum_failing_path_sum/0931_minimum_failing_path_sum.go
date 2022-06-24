package _931_minimum_failing_path_sum

import (
	"math"
)

// tptl. passed but slow
func minFallingPathSum(matrix [][]int) int {
	res := math.MaxInt64
	dp := make(map[[2]int]int)
	for i := range matrix {
		res = min(res, rec(dp, matrix, 0, i))
	}

	return res
}

func rec(dp map[[2]int]int, matrix [][]int, row, col int) int {
	if row == len(matrix) {
		return 0
	}

	if col < 0 || col == len(matrix[0]) {
		return math.MaxInt32
	}

	if res, ok := dp[[2]int{row, col}]; !ok {
		s1 := rec(dp, matrix, row+1, col-1)
		s2 := rec(dp, matrix, row+1, col)
		s3 := rec(dp, matrix, row+1, col+1)

		res = min3(s1, s2, s3) + matrix[row][col]

		dp[[2]int{row, col}] = res

		return res
	} else {
		return res
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func min3(a, b, c int) int {
	res := a
	if b < res {
		res = b
	}

	if c < res {
		res = c
	}

	return res
}
