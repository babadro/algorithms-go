package _2132_stamping_the_grid

// https://leetcode.com/problems/stamping-the-grid/discuss/1675412/JavaC%2B%2BPython-Calulate-the-sub-matrix-sum-twice
// passed. tptl. hard.
// todo 1 understand
func possibleToStamp2(M [][]int, h int, w int) bool {
	m, n := len(M), len(M[0])

	A, B, good := makeIntMatrix(m+1, n+1), makeIntMatrix(m+1, n+1), makeIntMatrix(m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			A[i+1][j+1] = A[i+1][j] + A[i][j+1] - A[i][j] + (1 - M[i][j])
			if i+1 >= h && j+1 >= w {
				x, y := i+1-h, j+1-w
				if A[i+1][j+1]-A[x][j+1]-A[i+1][y]+A[x][y] == w*h {
					good[i][j] = 1
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			B[i+1][j+1] = B[i+1][j] + B[i][j+1] - B[i][j] + good[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x, y := min(i+h, m), min(j+w, n)
			if M[i][j] == 0 && B[x][y]-B[i][y]-B[x][j]+B[i][j] == 0 {
				return false
			}
		}
	}

	return true

}

func makeIntMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	return matrix
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// faster solution. passed. not mine
func possibleToStamp3(grid [][]int, stampHeight, stampWidth int) bool {
	m, n := len(grid), len(grid[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	diff := make([][]int, m+1)
	diff[0] = make([]int, n+1)
	for i, row := range grid {
		sum[i+1] = make([]int, n+1)
		for j, v := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
		}
		diff[i+1] = make([]int, n+1)
	}
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				x, y := i+stampHeight, j+stampWidth
				if x <= m && y <= n && sum[x][y]-sum[x][j]-sum[i][y]+sum[i][j] == 0 {
					diff[i][j]++
					diff[i][y]--
					diff[x][j]--
					diff[x][y]++
				}
			}
		}
	}

	cnt := make([]int, n+1)
	pre := make([]int, n+1)
	for i, row := range grid {
		for j, v := range row {
			cnt[j+1] = cnt[j] + pre[j+1] - pre[j] + diff[i][j]
			if v == 0 && cnt[j+1] == 0 {
				return false
			}
		}
		cnt, pre = pre, cnt
	}
	return true
}
