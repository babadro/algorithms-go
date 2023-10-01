package _2088_count_fertile_piramyds_in_a_land

func countPyramids(grid [][]int) int {
	counter := 0

	for row := range grid {
		for col := range grid[0] {
			if y := row + 1; y < len(grid) {
				flag := true
				for x := col - 1; x < col+2; x++ {
					if x < 0 || x == len(grid[0]) || grid[y][x] == 0 {
						flag = false
						break
					}
				}

				if flag {
					counter++
				}
			}

			if y := row - 1; y >= 0 {
				flag := true
				for x := col - 1; x < col+2; x++ {
					if x < 0 || x == len(grid[0]) || grid[y][x] == 0 {
						flag = false
						break
					}
				}

				if flag {
					counter++
				}
			}
		}
	}

	return counter
}
