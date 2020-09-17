package _240_search_a_2d_matrix_ii

import "testing"

func Test_Search_a_2D_Matrix_II(t *testing.T) {
	testCsases := []struct {
		matrix [][]int
		target int
		result bool
	}{
		{
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			10,
			true,
		},
		{
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			20,
			false,
		},
		{
			[][]int{},
			20,
			false,
		},
	}

	for _, tc := range testCsases {
		r := searchMatrix3(tc.matrix, tc.target)
		if r != tc.result {
			t.Errorf("Source: %d \n Expected: %v \n Actual: %v",
				tc.matrix,
				tc.result,
				r)
		}
	}
}
