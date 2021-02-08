package _738_find_kth_largest_xor_coordinate_value

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	values := make([]int, 0, m*n)

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			val := 0
			for j := 0; j <= y; j++ {
				for k := 0; k <= x; k++ {
					val ^= matrix[j][k]
				}
			}

			values = append(values, val)
		}
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	return values[k-1]
}
