package _378_kth_smallest_element_in_sorted_matrix

import (
	"math"
	"sort"
)

// doesn't work
func kthSmallest2(matrix [][]int, k int) int {
	sqrt := math.Sqrt(float64(k))

	rounded := int(sqrt)
	if sqrt-float64(rounded) == 0 {
		return matrix[rounded-1][rounded-1]
	}

	n := rounded

	arr := make([]int, 0, n+n-1)

	for x := 0; x < len(matrix); x++ {
		arr = append(arr, matrix[n][x])
	}

	for y := 0; y < len(matrix)-1; y++ {
		arr = append(arr, matrix[y][n])
	}

	sort.Ints(arr)

	return arr[k-n*n-1]
}

// it works
func kthSmallestBruteForce(matrix [][]int, k int) int {
	arr := make([]int, 0, len(matrix)*len(matrix[0]))
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			arr = append(arr, matrix[y][x])
		}
	}

	sort.Ints(arr)
	return arr[k-1]
}
