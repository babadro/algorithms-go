package _054_spiral_matrix

// TODO 1
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[1])
	minSize := n
	if m < n {
		minSize = m
	}
	res := make([]int, 0, m*n)
	spiralCount := minSize/2 + minSize%2
	l, h := m, n
	x, y := 0, 0
	for i := 0; i < spiralCount; i++ {
		l, h = m-2*i, n-2*i
		for j := 0; j < l-1; j++ {
			res = append(res, matrix[y][x])
			x++
		}
		for j := 0; j < h-1; j++ {
			res = append(res, matrix[y][x])
			y++
		}
		for j := 0; j < l-1; j++ {
			res = append(res, matrix[y][x])
			x--
		}
		for j := 0; j < h-1; j++ {
			res = append(res, matrix[y][x])
			y--
		}
	}
	return res
}
