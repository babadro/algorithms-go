package _2713_maximum_strictly_increasing_cells_in_a_matrix

import "sort"

type Pos struct {
	y, x int
}

// #bnsrg
func maxIncreasingCells(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	horizontalMax, verticalMax := make([]int, n), make([]int, m)
	reversedMap := make(map[int][]Pos)
	cache := make([]int, m*n)
	var uniqValues []int

	for y := range mat {
		for x := range mat[0] {
			val := mat[y][x]
			if len(reversedMap[val]) == 0 {
				uniqValues = append(uniqValues, val)
			}

			reversedMap[val] = append(reversedMap[val], Pos{y, x})
		}
	}

	sort.Ints(uniqValues)

	ans := 0
	for _, val := range uniqValues {
		for i, pos := range reversedMap[val] {
			cache[i] = max(horizontalMax[pos.y], verticalMax[pos.x]) + 1
			ans = max(ans, cache[i])
		}

		for i, pos := range reversedMap[val] {
			horizontalMax[pos.y] = max(horizontalMax[pos.y], cache[i])
			verticalMax[pos.x] = max(verticalMax[pos.x], cache[i])
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
