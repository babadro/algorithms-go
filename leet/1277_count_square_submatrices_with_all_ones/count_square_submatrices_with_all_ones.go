package _1277_count_square_submatrices_with_all_ones

// bruteforce. passed. todo 1 look for better solution
func countSquares(matrix [][]int) int {
	counter, m, n := 0, len(matrix), len(matrix[0])
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			edge := n - x
			if height := m - y; height < edge {
				edge = height
			}
		Loop:
			for i := 0; i < edge; i++ {
				for j := 0; j <= i; j++ {
					if matrix[y+i][x+j] == 0 || matrix[y+j][x+i] == 0 {
						break Loop
					}
				}

				counter++
			}
		}
	}

	return counter
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// todo 1. just for example. need to understand
func countSquares2(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	count := 0
	//For first row
	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = matrix[0][i]
		count += dp[0][i]
	}
	//for first colomn
	for i := 1; i < len(matrix); i++ {
		dp[i][0] = matrix[i][0]
		count += dp[i][0]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = matrix[i][j] + min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1])
			}
			count += dp[i][j]
		}
	}
	return count
}
