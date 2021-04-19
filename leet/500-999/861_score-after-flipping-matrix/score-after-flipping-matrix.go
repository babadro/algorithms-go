package _861_score_after_flipping_matrix

// passed. tptl. medium. todo 1 look for greedy approach in solution - it's very short.
func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])

	for y := 0; y < len(A); y++ {
		if A[y][0] == 0 {
			inverseRow(A, y)
		}
	}

	for x := 1; x < n; x++ {
		count0, count1 := 0, 0
		for y := 0; y < m; y++ {
			if A[y][x] == 1 {
				count1++
			} else {
				count0++
			}
		}

		if count0 > count1 {
			inverseCol(A, x)
		}
	}

	score := 0
	for y := 0; y < m; y++ {
		num := 0
		for x := 0; x < n; x++ {
			num <<= 1
			num += A[y][x]
		}

		score += num
	}

	return score
}

func inverseRow(matrix [][]int, y int) {
	for x := range matrix[y] {
		if matrix[y][x] == 1 {
			matrix[y][x] = 0
		} else {
			matrix[y][x] = 1
		}
	}
}

func inverseCol(matrix [][]int, x int) {
	for y := 0; y < len(matrix); y++ {
		if matrix[y][x] == 1 {
			matrix[y][x] = 0
		} else {
			matrix[y][x] = 1
		}
	}
}
