package _2132_stamping_the_grid

// doesn't work.
func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	for y := range grid {
		for x := 0; x < len(grid[y]); x++ {
			start := x
			if grid[y][start] == 1 {
				continue
			}

			finish := start + 1
			for ; finish < len(grid[y]); finish++ {
				if grid[y][finish] == 1 {
					break
				}
			}

			if finish-start >= stampWidth {
				for k := start; k < finish; k++ {
					grid[y][k]--
				}
			}

			x = finish
		}
	}

	for x := range grid[0] {
		for y := 0; y < len(grid); y++ {
			start := y
			if grid[start][x] == 1 {
				continue
			}

			finish := start + 1
			for ; finish < len(grid); finish++ {
				if grid[finish][x] == 1 {
					break
				}
			}

			if finish-start >= stampHeight {
				for k := start; k < finish; k++ {
					grid[k][x]--
				}
			}

			y = finish
		}
	}

	for y := range grid {
		for x := range grid[y] {
			cell := grid[y][x]

			if cell == 1 {
				continue
			}

			if cell != -2 {
				return false
			}
		}
	}

	return true
}
