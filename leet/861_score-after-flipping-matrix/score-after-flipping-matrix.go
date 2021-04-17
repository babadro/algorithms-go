package _861_score_after_flipping_matrix

// todo 1
func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])

	for x := 0; x < n; x++ {
		count0, count1 := 0, 0
		for y := 0; y < m; y++ {
			if A[y][x] == 1 {
				count1++
			} else {
				count0++
			}

			if coun1
		}
	}
}
