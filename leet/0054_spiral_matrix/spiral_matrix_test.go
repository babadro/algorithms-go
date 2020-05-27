package _054_spiral_matrix

import "testing"

var input1 = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}
var input2 = [][]int{
	{5, 1, 9, 11},
	{2, 4, 8, 10},
	{13, 3, 6, 7},
	{15, 14, 12, 16},
}

func TestSpiralOrder(t *testing.T) {
	t.Log(spiralOrder(input1))
}
