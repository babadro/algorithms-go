package matrix

var _1x1 = [][]int{
	{1},
}

var _2x2 = [][]int{
	{1, 2},
	{3, 4},
}

var _3x3 = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}
var _4x4 = [][]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
	{13, 14, 15, 16},
}

var _4x1 = [][]int{
	{1, 2, 3, 4},
}

var _1x4 = [][]int{
	{1},
	{2},
	{3},
	{4},
}

func New(n, m int) [][]int {
	matrix := make([][]int, m)
	i := 1
	for j := range matrix {
		matrix[j] = make([]int, n)
		for k := range matrix[j] {
			matrix[j][k] = i
			i++
		}
	}
	return matrix
}
