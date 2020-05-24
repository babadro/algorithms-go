package _048_rotate_image

// TODO 1
func rotate(matrix [][]int) {
	matrixSize := len(matrix)
	for y := 0; y < matrixSize/2; y++ {
		rowLen := matrixSize - 2*y
		for x := 0; x < rowLen-1; x++ {
			move(matrix, x, y)
		}
	}
}

func move(matrix [][]int, x, y int) {
	step, matrixSize := 3, 4
	tmp, newX, newY := 0, 0, 0
	matrix[y]
}
