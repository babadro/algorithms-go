package _931_minimum_failing_path_sum

import (
	"math"
)

// tptl.  bottom up
func minFallingPathSum2(matrix [][]int) int {
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len(matrix))
	}

	for row := len(matrix) - 1; row >= 0; row-- {
		for col := 0; col < len(matrix[0]); col++ {
			var minRes int
			if len(matrix[0]) == 1 {
				minRes = dp[1][col]
			} else if col == 0 {
				minRes = min(dp[1][col], dp[1][col+1])
			} else if col == len(matrix[0])-1 {
				minRes = min(dp[1][col-1], dp[1][col])
			} else {
				minRes = min3(dp[1][col-1], dp[1][col], dp[1][col+1])
			}

			dp[0][col] = minRes + matrix[row][col]
		}

		dp[0], dp[1] = dp[1], dp[0]
	}

	res := math.MaxInt64
	for _, num := range dp[1] {
		res = min(res, num)
	}

	return res
}

// tptl. top down. passed but slow
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
