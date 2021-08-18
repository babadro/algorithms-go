package contest

import (
	"math"
	"sort"
)

// passed. #hard #tptl mine
// todo 2 find shorter solution
func latestDayToCross(row int, col int, cells [][]int) int {
	matrix := make([][]int, row)
	for i := range matrix {
		matrix[i] = make([]int, col)
	}

	for i, cell := range cells {
		y, x := cell[0]-1, cell[1]-1
		matrix[y][x] = i + 1
	}

	lastDayIDX := sort.Search(len(cells), func(i int) bool {
		found := false
		for x := 0; x < len(matrix[0]); x++ {
			if crawl(matrix, 0, x, i) {
				found = true
				break
			}
		}

		clean(matrix)

		return !found
	})

	return lastDayIDX
}

func crawl(matrix [][]int, y, x, idx int) bool {
	if y < 0 || y == len(matrix) || x < 0 || x == len(matrix[0]) {
		return false
	}

	if flooded(matrix, y, x, idx) || visited(matrix, y, x) {
		return false
	}

	if val := matrix[y][x]; val == 0 {
		matrix[y][x] = math.MaxInt64
	} else {
		matrix[y][x] = -val
	}

	if y == len(matrix)-1 {
		return true
	}

	return crawl(matrix, y+1, x, idx) ||
		crawl(matrix, y, x+1, idx) ||
		crawl(matrix, y, x-1, idx) ||
		crawl(matrix, y-1, x, idx)
}

func clean(matrix [][]int) {
	for _, row := range matrix {
		for i := range row {
			if row[i] == math.MaxInt64 {
				row[i] = 0
			} else if row[i] < 0 {
				row[i] = -row[i]
			}
		}
	}
}

func visited(matrix [][]int, y, x int) bool {
	return matrix[y][x] == math.MaxInt64 || matrix[y][x] < 0
}

func flooded(matrix [][]int, y, x, idx int) bool {
	return matrix[y][x] > 0 && matrix[y][x] <= idx+1
}
