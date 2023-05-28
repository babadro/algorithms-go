package _1254_number_of_closed_islands

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
		return true
	}

	if grid[y][x] != 0 {
		return true
	}

	grid[y][x] = 2

	nearTheBorder := y == 0 || y == len(grid)-1 || x == 0 || x == len(grid[y])-1

	res := !nearTheBorder

	if !crawl(x-1, y, grid) {
		res = false
	}

	if !crawl(x+1, y, grid) {
		res = false
	}

	if !crawl(x, y-1, grid) {
		res = false
	}

	if !crawl(x, y+1, grid) {
		res = false
	}

	return res
}
