package _1727_largest_submatrix_with_rearrangements

import "sort"

// passed. tptl.
func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	for x := 0; x < n; x++ {
		counter := 0
		for y := 0; y < m; y++ {
			if matrix[y][x] == 0 {
				counter = 0
			} else {
				counter++
				matrix[y][x] = counter
			}
		}
	}

	max := 0
	for y := 0; y < m; y++ {
		sort.Slice(matrix[y], func(i, j int) bool {
			return matrix[y][i] > matrix[y][j]
		})

		maxRow := 0
		for x := 0; x < n; x++ {
			if matrix[y][x] == 0 {
				break
			}

			res := matrix[y][x] * (x + 1)
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
