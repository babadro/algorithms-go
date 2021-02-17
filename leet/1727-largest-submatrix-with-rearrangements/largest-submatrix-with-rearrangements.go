package _1727_largest_submatrix_with_rearrangements

import "sort"

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	rows := make([][]int, m)
	for i := 0; i < m; i++ {
		rows[i] = make([]int, n)
	}

	max := 0
	for y := 0; y < m; y++ {
		counter := 0
		for x := 0; x < n; x++ {
			if matrix[y][x] == 1 {
				counter++
				rows[y][x] = counter
			}
		}

		sort.Slice(rows[y], func(i, j int) bool {
			return rows[y][i] > rows[y][j]
		})

		maxRow := 0
		for x := 0; x < n; x++ {
			if rows[y][x] == 0 {
				break
			}

			res := rows[y][x] * (x + 1)
			if res > maxRow {
				maxRow = res
			}
		}

		if maxRow > max {
			max = maxRow
		}
	}

	return max
}
