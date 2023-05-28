package _1254_number_of_closed_islands

// #bnsrg
func closedIsland(grid [][]int) int {
	counter := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 && crawl(x, y, grid) {
				counter++
			}
		}
	}

	return counter
}

func crawl(x, y int, grid [][]int) bool {
	if y < 0 || y == len(grid) || x < 0 || x == len(grid[y]) {
		return false
	}

	if grid[y][x] != 0 {
		return true
	}

	grid[y][x] = 2

	res1 := crawl(x+1, y, grid)
	res2 := crawl(x-1, y, grid)
	res3 := crawl(x, y+1, grid)
	res4 := crawl(x, y-1, grid)

	return res1 && res2 && res3 && res4
}
