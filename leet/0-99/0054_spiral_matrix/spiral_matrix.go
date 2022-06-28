package _054_spiral_matrix

// passed. best solution
func spiralOrder2(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	total := n * m
	up, down, left, right := 0, n-1, 0, m-1
	res := make([]int, 0, total)
	for len(res) < total {
		for x := left; x <= right && len(res) < total; x++ {
			res = append(res, matrix[up][x])
		}

		for y := up + 1; y <= down-1 && len(res) < total; y++ {
			res = append(res, matrix[y][right])
		}

		for x := right; x >= left && len(res) < total; x-- {
			res = append(res, matrix[down][x])
		}

		for y := down - 1; y >= up+1 && len(res) < total; y-- {
			res = append(res, matrix[y][left])
		}

		left++
		right--
		up++
		down--
	}

	return res
}

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
