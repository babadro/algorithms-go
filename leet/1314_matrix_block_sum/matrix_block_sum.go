package _1314_matrix_block_sum

func matrixBlockSum2(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])

	answer, sumMat := make([][]int, m), make([][]int, m)
	for i := range answer {
		answer[i] = make([]int, n)
		sumMat[i] = make([]int, n)
	}

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			sumMat[y][x] = mat[y][x]
		}
	}

	helper(sumMat)

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			y1 := y + K
			if y1 >= m {
				y1 = m - 1
			}

			x1 := x + K
			if x1 >= n {
				x1 = n - 1
			}

			res := sumMat[y1][x1]

			y2 := y - K - 1
			if y2 >= 0 {
				res -= sumMat[y2][x1]
			}

			x2 := x - K - 1
			if x2 >= 0 {
				res -= sumMat[y1][x2]
			}

			if y2 >= 0 && x2 >= 0 {
				res += sumMat[y2][x2]
			}

			answer[y][x] = res
		}
	}

	return answer
}

func helper(mat [][]int) {
	m, n := len(mat), len(mat[0])

	rowSum := make([]int, n)
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			rowSum[x] = mat[y][x]
			if prevX := x - 1; prevX >= 0 {
				rowSum[x] += rowSum[prevX]
				mat[y][x] += rowSum[prevX]
			}

			if prevY := y - 1; prevY >= 0 {
				mat[y][x] += mat[prevY][x]
			}
		}
	}
}

// brutforce. passed. too slow.
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
