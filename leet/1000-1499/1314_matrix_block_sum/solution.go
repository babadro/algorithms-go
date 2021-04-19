package _1314_matrix_block_sum

// best solution. passed. tptl
func matrixBlockSum3(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])
	sumMat := make([][]int, m+1)
	res := make([][]int, m)

	for i := range res {
		res[i] = make([]int, n)
	}

	for i := range sumMat {
		sumMat[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sumMat[i+1][j+1] = sumMat[i+1][j] + sumMat[i][j+1] - sumMat[i][j] + mat[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r1, c1, r0, c0 := min(m, i+K+1), min(n, j+K+1), max(0, i-K), max(0, j-K)
			res[i][j] = sumMat[r1][c1] - sumMat[r1][c0] - sumMat[r0][c1] + sumMat[r0][c0]
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// just for understanding the result...
func helper2(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	summedAreaTable := make([][]int, m+1)
	for i := range summedAreaTable {
		summedAreaTable[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			summedAreaTable[i+1][j+1] = summedAreaTable[i+1][j] + summedAreaTable[i][j+1] - summedAreaTable[i][j] + mat[i][j]
		}
	}

	return summedAreaTable
}
