package _240_search_a_2d_matrix_ii

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	leftY, leftX, rightY, rightX := 0, 0, len(matrix), len(matrix[0])

	for leftY < rightY && leftX < rightX {
		middleY, middleX := (leftY+rightY)/2, (leftX+rightX)/2
		if num := matrix[middleY][middleX]; num > target {
			rightX, rightY = middleX, middleY
			continue
		} else if num < target {
			leftX, leftY = middleX+1, middleY+1
			continue
		}
		return true
	}

}

// brutforce. status: accepted
func searchMatrixBrutforce(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if matrix[y][x] == target {
				return true
			}
		}
	}

	return false
}
