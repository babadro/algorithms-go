package _054_spiral_matrix

// 100% 100%
// TODO 3 it's ok, but the shortest examples exist
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	minSize := n
	if m < n {
		minSize = m
	}
	totalCount := m * n
	res := make([]int, 0, totalCount)
	spiralCount := minSize/2 + minSize%2
	l, h := m, n
	x, y := 0, 0
	for i := 0; i < spiralCount; i++ {
		l, h = n-2*i, m-2*i
		for j := 0; j < l-1; j++ {
			res = append(res, matrix[y][x])
			x++
		}
		for j := 0; j < h-1; j++ {
			res = append(res, matrix[y][x])
			y++
		}
		for j := 0; j < l-1 && len(res) < totalCount; j++ {
			res = append(res, matrix[y][x])
			x--
		}
		for j := 0; j < h-2 && len(res) < totalCount; j++ {
			res = append(res, matrix[y][x])
			y--
		}
		if len(res) < totalCount {
			res = append(res, matrix[y][x])
			x++
		}
	}
	return res
}
