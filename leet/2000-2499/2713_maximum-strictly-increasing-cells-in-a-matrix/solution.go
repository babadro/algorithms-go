package _2713_maximum_strictly_increasing_cells_in_a_matrix

import "sort"

type Pos struct {
	i, j int
}

func maxIncreasingCells(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	horizontalMax, verticalMax := make([]int, n), make([]int, m)
	reverseMat := map[int][]Pos{}
	uniqueVals := []int{}
	cache := make([]int, n*m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			val := mat[i][j]
			if len(reverseMat[val]) == 0 {
				uniqueVals = append(uniqueVals, val)
			}
			reverseMat[val] = append(reverseMat[val], Pos{i, j})
		}
	}

	sort.Ints(uniqueVals)
	ans := 0

	for _, val := range uniqueVals {
		for k, pos := range reverseMat[val] {
			cache[k] = max(horizontalMax[pos.i], verticalMax[pos.j]) + 1
			ans = max(ans, cache[k])
		}
		for k, pos := range reverseMat[val] {
			horizontalMax[pos.i] = max(horizontalMax[pos.i], cache[k])
			verticalMax[pos.j] = max(verticalMax[pos.j], cache[k])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
