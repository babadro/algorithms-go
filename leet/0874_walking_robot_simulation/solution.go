package _874_walking_robot_simulation

// passed. best solution. tptl
func robotSim2(commands []int, obstacles [][]int) int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m := make(map[[2]int]bool)
	for _, obstacle := range obstacles {
		m[[2]int{obstacle[0], obstacle[1]}] = true
	}

	var res, x, y, i int
	for _, cmd := range commands {
		switch cmd {
		case -1:
			i = (i + 1) % 4
		case -2:
			i = (i + 3) % 4
		default:
			for ; cmd > 0; cmd-- {
				if m[[2]int{x + dirs[i][0], y + dirs[i][1]}] {
					break
				}
				x, y = x+dirs[i][0], y+dirs[i][1]
			}

			res = max(res, x*x+y*y)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
