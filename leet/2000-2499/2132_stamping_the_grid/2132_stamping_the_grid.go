package _2132_stamping_the_grid

// todo 1 - doesn't work. Fix last test case
// #hard
func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m := make(map[[2]int][2]bool)

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
					key := [2]int{k, y}
					v := m[key]
					v[0] = true
					m[key] = v
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
					key := [2]int{x, k}
					v := m[key]
					v[1] = true
					m[key] = v
				}
			}

			y = finish
		}
	}

	for _, v := range m {
		if !v[0] || !v[1] {
			return false
		}
	}

	return true
}
