package _200_number_of_islands

type point struct {
	x, y int
}

func numIslands(grid [][]byte) int {
	//islandToPoint := make(map[int]point)
	pointToIsland := make(map[point]int)

	islandNum := 0
	for x, arr := range grid {
		for y, val := range arr {
			if val == 0 {
				continue
			}
			currPoint := point{x, y}
			if y > 0 && arr[y-1] != 0 {
				adjacentPoint := point{x, y - 1}
				pointToIsland[currPoint] = pointToIsland[adjacentPoint]
				continue
			}
			islandNum++
			pointToIsland[currPoint] = islandNum
			if x > 0 && grid[x-1][y] != 0 {
				adjacentPoint := point{x - 1, y}
				if pointToIsland[adjacentPoint] != pointToIsland[currPoint] {
					pointToIsland[adjacentPoint] = pointToIsland[currPoint]
				}
			} else {
				islandNum++
				pointToIsland[point{x, y}] = islandNum
			}
		}
	}

	return islandNum
}
