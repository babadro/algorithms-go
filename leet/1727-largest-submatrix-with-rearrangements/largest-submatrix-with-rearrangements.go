package _1727_largest_submatrix_with_rearrangements

import "sort"

// passed. tptl.
// contest 224
func largestSubmatrix(matrix [][]int) int {
	for y := 1; y < len(matrix); y++ {
		for x := range matrix[0] {
			if matrix[y][x] == 1 {
				matrix[y][x] += matrix[y-1][x]
			}
		}
	}

	n := len(matrix[0])
	max := 0
	for _, row := range matrix {
		sort.Ints(row)
		for i := 1; i <= n; i++ {
			if res := row[n-i] * i; res > max {
				max = res
			}
		}
	}

	return max
}
