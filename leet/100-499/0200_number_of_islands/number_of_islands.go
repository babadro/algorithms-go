package _200_number_of_islands

func numIslands(grid [][]byte) int {
	lenY := len(grid)
	if lenY == 0 {
		return 0
	}
	lenX := len(grid[0])
	counter := 0
	for y, arr := range grid {
		for x := range arr {
			if grid[y][x] == '1' {
				counter++
				walk(x, y, lenY, lenX, grid)
			}
		}
	}
	return counter
}

func walk(x, y, lenY, lenX int, grid [][]byte) {
	if x < 0 || x == lenX || y < 0 || y == lenY || grid[y][x] == '0' {
		return
	}
	grid[y][x] = '0'
	walk(x, y-1, lenY, lenX, grid)
	walk(x, y+1, lenY, lenX, grid)
	walk(x-1, y, lenY, lenX, grid)
	walk(x+1, y, lenY, lenX, grid)
}
