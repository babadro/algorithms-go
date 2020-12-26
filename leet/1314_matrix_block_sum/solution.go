package _1314_matrix_block_sum

// best solution. todo 1 understand it
func matrixBlockSum3(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])
	summedAreaTable := make([][]int, m+1)
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := range summedAreaTable {
		summedAreaTable[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			summedAreaTable[i+1][j+1] = summedAreaTable[i+1][j] + summedAreaTable[i][j+1] - summedAreaTable[i][j] + mat[i][j]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Try to understand why r0, c0 = i - K, j - K. but r1, c1 = i + K + 1, j + K + 1.
			r1, c1, r0, c0 := min(m, i+K+1), min(n, j+K+1), max(0, i-K), max(0, j-K)
			ans[i][j] = summedAreaTable[r1][c1] - summedAreaTable[r1][c0] - summedAreaTable[r0][c1] + summedAreaTable[r0][c0]
		}
	}
	return ans
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
