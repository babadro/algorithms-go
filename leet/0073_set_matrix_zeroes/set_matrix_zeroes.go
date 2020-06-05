package _073_set_matrix_zeroes

func setZeros(matrix [][]int) {
	row, col, zeroFound := 0, 0, false
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == 0 {
				if !zeroFound {
					zeroFound = true
					row, col = y, x
				}
				matrix[row][x] = 0
				matrix[y][col] = 0
			}
		}
	}

	if !zeroFound {
		return
	}

	for x := range matrix[row] {
		if matrix[row][x] == 0 && x != col {
			for y := 0; y < len(matrix); y++ {
				matrix[y][x] = 0
			}
		}
		matrix[row][x] = 0
	}

	for y := 0; y < len(matrix); y++ {
		if matrix[y][col] == 0 {
			for x := 0; x < len(matrix[y]); x++ {
				matrix[y][x] = 0
			}
		}
		matrix[y][col] = 0
	}
}
