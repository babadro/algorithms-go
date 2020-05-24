package _048_rotate_image

// 100% 100%
// TODO 2 need to understand rotate2 - it's faster than mine solution 6x times.
func rotate(matrix [][]int) {
	matrixSize := len(matrix)
	newY, newX, x, y := 0, 0, 0, 0
	for i := 0; i < matrixSize/2; i++ {
		step := matrixSize - 2*i - 1
		for j := 0; j < step; j++ {
			y, x = i, i+j
			var tmp int

			// edge 1 => edge 2
			newY, newX = y, i+step
			if j > 0 {
				newY = i + (j+step)%step
			}
			tmp, matrix[newY][newX] = matrix[newY][newX], matrix[y][x]
			y, x = newY, newX

			// edge 2 => edge 3
			newY, newX = i+step, x
			if j > 0 {
				newX = i + step - (j+step)%step
			}
			tmp, matrix[newY][newX] = matrix[newY][newX], tmp
			y, x = newY, newX

			// edge 3 => edge 4
			newY, newX = y, i
			if j > 0 {
				newY = i + step - (j+step)%step
			}
			tmp, matrix[newY][newX] = matrix[newY][newX], tmp
			y, x = newY, newX

			// edge 4 => edge 1
			newY, newX = i, x
			if j > 0 {
				newX = i + (j+step)%step
			}
			matrix[newY][newX] = tmp
		}
	}
}
