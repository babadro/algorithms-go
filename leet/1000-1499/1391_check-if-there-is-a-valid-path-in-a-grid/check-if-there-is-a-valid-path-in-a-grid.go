package _1391_check_if_there_is_a_valid_path_in_a_grid

// tptl. passed. medium (but for me it's almost hard)
// todo 2 look for faster solution, probably iterative, if possible
func hasValidPath(grid [][]int) bool {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return true
	}

	next1, next2 := adjacentCells(0, 0, grid)
	visited1, visited2 := map[[2]int]bool{}, map[[2]int]bool{}
	start := [2]int{0, 0}
	visited1[start], visited2[start] = true, true

	return helper(start, next1, grid, visited1) || helper(start, next2, grid, visited2)
}

func helper(prev, curr [2]int, grid [][]int, visited map[[2]int]bool) bool {
	m, n := len(grid), len(grid[0])
	x, y := curr[0], curr[1]
	if x < 0 || x == n || y < 0 || y == m {
		return false
	}

	if !match(prev, curr, grid) {
		return false
	}

	if x == n-1 && y == m-1 {
		return true
	}

	if visited[curr] {
		return false
	}

	visited[curr] = true

	next1, next2 := adjacentCells(x, y, grid)

	return helper(curr, next1, grid, visited) || helper(curr, next2, grid, visited)
}

func adjacentCells(x, y int, grid [][]int) ([2]int, [2]int) {
	street := grid[y][x]
	switch street {
	case 1:
		return [2]int{x - 1, y}, [2]int{x + 1, y}
	case 2:
		return [2]int{x, y - 1}, [2]int{x, y + 1}
	case 3:
		return [2]int{x - 1, y}, [2]int{x, y + 1}
	case 4:
		return [2]int{x, y + 1}, [2]int{x + 1, y}
	case 5:
		return [2]int{x - 1, y}, [2]int{x, y - 1}
	case 6:
		return [2]int{x, y - 1}, [2]int{x + 1, y}
	}

	return [2]int{}, [2]int{}
}

func match(prev, next [2]int, grid [][]int) bool {
	x, y := next[0], next[1]
	adj1, adj2 := adjacentCells(x, y, grid)

	return adj1 == prev || adj2 == prev
}
