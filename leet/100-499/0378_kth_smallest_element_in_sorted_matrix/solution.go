package _378_kth_smallest_element_in_sorted_matrix

// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/discuss/352566/Simple-binary-Search-in-Go-20-ms-beat-100
// tiq. todo 2 need to remember
func kthSmallestBestSolution(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]

	for l < r {
		m := (l + r) / 2
		cnt := 0
		j := n - 1
		for i := 0; i < n; i++ {
			for j >= 0 && matrix[i][j] > m {
				j -= 1
			}
			cnt += j + 1
		}

		if cnt < k {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}
