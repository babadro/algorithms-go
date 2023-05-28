package _1254_number_of_closed_islands

// #bnsrg
func closedIsland(grid [][]int) int {
	counter, ok := 0, false
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				ok = true
				crawl(x, y, grid, &ok)

				if ok {
					counter++
				}
			}
		}
	}

	return counter
}

func crawl(x, y int, grid [][]int, ok *bool) {
	if y < 0 || y == len(grid) || x < 0 || x == len(grid[y]) || grid[y][x] != 0 {
		return
	}

	grid[y][x] = 2

	if y == 0 || y == len(grid)-1 || x == 0 || x == len(grid[y])-1 {
		*ok = false
	}

	crawl(x-1, y, grid, ok)
	crawl(x+1, y, grid, ok)
	crawl(x, y-1, grid, ok)
	crawl(x, y+1, grid, ok)
}
