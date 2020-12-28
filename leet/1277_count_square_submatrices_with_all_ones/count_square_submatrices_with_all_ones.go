package _1277_count_square_submatrices_with_all_ones

// passed. not mine. best solution. tptl
func countSquares2(matrix [][]int) int {
	counter := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] > 0 && i > 0 && j > 0 {
				matrix[i][j] = min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1])) + 1
			}

			counter += matrix[i][j]
		}
	}

	return counter
}

// bruteforce. passed.
func countSquares(matrix [][]int) int {
	counter, m, n := 0, len(matrix), len(matrix[0])
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			edge := n - x
			if height := m - y; height < edge {
				edge = height
			}
		Loop:
			for i := 0; i < edge; i++ {
				for j := 0; j <= i; j++ {
					if matrix[y+i][x+j] == 0 || matrix[y+j][x+i] == 0 {
						break Loop
					}
				}

				counter++
			}
		}
	}

	return counter
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
