package _934_shortest_bridge

import "math"

// bnsrg
func shortestBridge(grid [][]int) int {
	flag := false
	arr := make([][2]int, 0, 100)
	minDist := math.MaxInt64
	for y := range grid {
		for x := range grid[y] {
			if val := grid[y][x]; val == 1 {
				if !flag {
					flag = true
					crawl(grid, &arr, y, x)

					continue
				}

				for _, coord := range arr {
					y2, x2 := coord[0], coord[1]

					distY, distX := abs(y)-y2, abs(x)-x2

					if distY == 0 {
						minDist = min(minDist, distX-1)
					} else if distX == 0 {
						minDist = min(minDist, distY-1)
					} else if distY == 1 && distX == 1 {
						return 1
					} else {
						minDist = min(minDist, distY+distX-1)
					}
				}

			}
		}
	}

	return minDist
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func crawl(grid [][]int, arr *[][2]int, y, x int) {
	if y < 0 || y == len(grid) || x < 0 || x == len(grid[0]) {
		return
	}

	if grid[y][x] != 1 {
		return
	}

	grid[y][x] = 2
	*arr = append(*arr, [2]int{y, x})

	crawl(grid, arr, y+1, x)
	crawl(grid, arr, y-1, x)
	crawl(grid, arr, y, x+1)
	crawl(grid, arr, y, x-1)
}
