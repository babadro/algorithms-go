package _1314_matrix_block_sum

// brutforce. passed. too slow. todo 1 write more effective solution
func matrixBlockSum(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])

	answer := make([][]int, m)
	for i := range answer {
		answer[i] = make([]int, n)
	}

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			sum := 0
			for j := y - K; j <= y+K; j++ {
				for k := x - K; k <= x+K; k++ {
					if j < 0 || j >= m || k < 0 || k >= n {
						continue
					}

					sum += mat[j][k]
				}
			}

			answer[y][x] = sum
		}
	}

	return answer
}
