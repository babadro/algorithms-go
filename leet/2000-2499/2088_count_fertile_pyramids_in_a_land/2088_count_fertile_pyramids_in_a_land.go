package _2088_count_fertile_piramyds_in_a_land

func countPyramids(grid [][]int) int {
	counter := 0

	for row := range grid {
		for col := range grid[0] {

		Loop1:
			for rowLen, y, startX := 3, row+1, col-1; ; rowLen, y, startX = rowLen+2, y+1, startX-1 {
				for x := startX; x < rowLen+startX; x++ {
					if y == len(grid) || x < 0 || x == len(grid[0]) || grid[y][x] == 0 {
						break Loop1
					}
				}

				counter++
			}

		Loop2:
			for rowLen, y, startX := 3, row-1, col-1; ; rowLen, y, startX = rowLen+2, y-1, startX-1 {
				for x := startX; x < rowLen+startX; x++ {
					if y < 0 || x < 0 || x == len(grid[0]) || grid[y][x] == 0 {
						break Loop2
					}
				}

				counter++
			}
		}
	}

	return counter
}
