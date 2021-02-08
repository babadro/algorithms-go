package _1738_find_kth_largest_xor_coordinate_value

import "sort"

// passed, but slow.
func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	values := make([]int, 0, m*n)

	xors := make([][]int, m)
	for i := range xors {
		xors[i] = make([]int, n)
	}

	for y := 0; y < m; y++ {
		val := 0
		for x := 0; x < n; x++ {
			val ^= matrix[y][x]
			xors[y][x] = val
		}
	}

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			val := 0
			for j := 0; j <= y; j++ {
				val ^= xors[j][x]
			}

			values = append(values, val)
		}
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	return values[k-1]
}
