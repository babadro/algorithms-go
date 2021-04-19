package _976_largest_perimeter_triangel

import "sort"

// passed. easy. tptl
func largestPerimeter(A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})

	for i := 2; i < len(A); i++ {
		if A[i]+A[i-1] > A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}

	return 0
}
