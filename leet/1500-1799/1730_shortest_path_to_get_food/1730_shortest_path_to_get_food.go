package _1730_shortest_path_to_get_food

import "math"

// #bnsrg medium tle
func getFood(grid [][]byte) int {
	res := math.MaxInt64

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '*' {
				traverse(y, x, 0, grid, make(map[[2]int]int), &res)

				if res == math.MaxInt64 {
					return -1
				}

				return res
			}
		}
	}

	return 0
}

func traverse(y, x, counter int, grid [][]byte, visited map[[2]int]int, res *int) {
	key := [2]int{y, x}

	if y < 0 || y == len(grid) ||
		x < 0 || x == len(grid[y]) ||
		grid[y][x] == 'X' || *res <= counter {
		return
	}

	if cnt, ok := visited[key]; ok && cnt <= counter {
		return
	}

	if grid[y][x] == '#' {
		*res = min(*res, counter)

		return
	}

	visited[key] = counter

	traverse(y-1, x, counter+1, grid, visited, res)
	traverse(y+1, x, counter+1, grid, visited, res)
	traverse(y, x-1, counter+1, grid, visited, res)
	traverse(y, x+1, counter+1, grid, visited, res)
}
