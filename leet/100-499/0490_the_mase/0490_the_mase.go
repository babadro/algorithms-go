package _490_the_mase

var directions = [][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

// #bnsrg medium. passed but slow - 11%, 5,5%
// todo 2 find shorter and faster solution
func hasPath(maze [][]int, start []int, destination []int) bool {
	var res bool

	y, x := start[0], start[1]
	visited := make(map[[3]int]bool)

	for i, d := range directions {
		rec(maze, destination, y+d[0], x+d[1], i, &res, visited)
	}

	return res
}

func rec(maze [][]int, destination []int, y, x int, dir int, res *bool, visited map[[3]int]bool) bool {
	if y < 0 || y == len(maze) || x < 0 || x == len(maze[y]) ||
		maze[y][x] == 1 {
		return false
	}

	key := [3]int{y, x, dir}

	if visited[key] {
		return true
	}

	visited[key] = true

	ok := rec(maze, destination, y+directions[dir][0], x+directions[dir][1], dir, res, visited)

	if ok {
		return true
	}

	if destination[0] == y && destination[1] == x {
		*res = true

		return true
	}

	for i, d := range directions {
		rec(maze, destination, y+d[0], x+d[1], i, res, visited)
	}

	return true
}
