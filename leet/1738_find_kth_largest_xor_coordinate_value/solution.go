package _1738_find_kth_largest_xor_coordinate_value

import "sort"

// best solution. tptl. passed.
func kthLargestValue2(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	xor := make([]int, m*n)
	fill(xor, m, n, matrix)
	sort.Ints(xor)
	return xor[m*n-k]
}

// https://accu.org/journals/overload/20/109/lewin_1915/
// remember properties of XOR:
// a^a = 0; (a^b)^c = a^(b^c); a^b = b^a; 0^a = a; => a^a^a = a
func fill(xor []int, m, n int, matrix [][]int) {
	xor[0] = matrix[0][0]
	for i := 1; i < m; i++ {
		xor[i*n] = xor[(i-1)*n] ^ matrix[i][0]
	}

	for i := 1; i < n; i++ {
		xor[i] = xor[i-1] ^ matrix[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			idx := i*n + j
			xor[idx] = xor[idx-1] ^ xor[idx-n] ^ xor[idx-n-1] ^ matrix[i][j]
		}
	}
}
