package hasValidPath

type point struct {
	i, j int
}

func hasValidPath(grid [][]int) bool {
	roadMap := make(map[int][3]int)
	roadMap[1] = [3]int{1, 3, 5}
	roadMap[2] = [3]int{2, 5, 6}
	roadMap[3] = [3]int{5, 6, 2}
	roadMap[4] = [3]int{5, 6, 2}
	roadMap[6] = [3]int{}
	i, j, ok := 0, 0, true
	for {
		if i, j, ok = move(i, j, grid,); !ok {
			break
		}
	}
	return i == len(grid) && j == len(grid[i])
}

func move(i, j int, grid [][]int, visited map[point]bool) (int, int, bool) {
	in := grid[i][j]
	if in == 1 {
		newI, newJ := i, j+1
		if !visited[point{newI, newJ}] && newJ < len(grid[i]) &&
	}
}

func match(in, out int) bool {
	if in == 1 {
		return out == 1 || out == 3 || out == 5
	}
	if in == 2 {
		return out == 2 || out == 5 || out == 6
	}
	if in == 3 {
		return out ==
	}
}
