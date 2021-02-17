package _1727_largest_submatrix_with_rearrangements

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	rows, cols := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		rows[i], cols[i] = make([]int, n), make([]int, n)
	}

	for y := 0; y < m; y++ {
		counter := 0
		for x := 0; x < n; x++ {
			if matrix[y][x] == 1 {
				counter++
				rows[y][x] = counter
			}
		}
	}

	for x := 0; x < n; x++ {
		counter := 0
		for y := 0; y < m; y++ {
			if matrix[y][x] == 1 {
				counter++
				cols[y][x] = counter
			}
		}
	}

	max := 0
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			res := rows[y][x] * cols[y][x]
			if res > max {
				max = res
			}
		}
	}

	return max
}
