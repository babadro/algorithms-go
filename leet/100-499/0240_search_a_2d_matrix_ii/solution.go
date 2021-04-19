package _240_search_a_2d_matrix_ii

// best solution
func searchMatrix3(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		val := matrix[row][col]
		if val > target {
			col--
		} else if val < target {
			row++
		} else {
			return true
		}
	}

	return false
}

/*
// https://leetcode.com/problems/search-a-2d-matrix-ii/discuss/738691/Go-Binary-Search-(O-(N-*-log-N)
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) > 0 {
		rowsCount := len(matrix) - 1
		colsCount := len(matrix[0]) - 1

		for y := rowsCount; y >= 0; y-- {
			left := 0
			right := colsCount
			for left <= right {
				x := (right + left) / 2
				if matrix[y][x] == target {
					return true
				}
				if matrix[y][x] > target {
					right = x - 1
				} else {
					left = x + 1
				}
			}
		}
	}

	return false
}
*/
