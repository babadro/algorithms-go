package _490_the_mase

var directions = [][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func hasPath(maze [][]int, start []int, destination []int) bool {
	var res bool

	y, x := start[0], start[1]
	visited := make(map[[3]int]bool)

	for i := range directions {
		rec(maze, destination, y, x, i, &res, visited)
	}

	return res
}

func rec(maze [][]int, destination []int, curY, curX int, dir int, res *bool, visited map[[3]int]bool) {
	y, x := directions[dir][0], directions[dir][1]

	key := [3]int{y, x, dir}

	if visited[key] {
		return
	}

	if y == len(maze) || y < 0 || x == len(maze[y]) || x < 0 ||
		maze[y][x] == 1 {
		if curY == destination[0] && curX == destination[1] {
			*res = true

			return
		}

		for i := range directions {
			if i == dir {
				continue
			}

			rec(maze, destination, curY, curX, i, res, visited)
		}

		return
	}

	visited[key] = true

	rec(maze, destination, y, x, dir, res, visited)
}
