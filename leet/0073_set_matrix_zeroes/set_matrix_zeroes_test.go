package _073_set_matrix_zeroes

import "testing"

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0, 5},
		{4, 3, 1, 4},
		{0, 1, 1, 4},
		{1, 2, 1, 3},
		{0, 0, 1, 1},
	}
	_ = matrix
	matrix2 := [][]int{
		{1, 2, 3},
		{1, 0, 3},
		{1, 2, 3},
	}
	setZeros(matrix2)
	t.Log(matrix2)
}
