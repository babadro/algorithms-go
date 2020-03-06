package _200_number_of_islands

import "sort"

type point struct {
	x, y int
}

func numIslands(grid [][]byte) int {
	pointToIsland := make(map[point]int)

	islandId := 0
	for x, arr := range grid {
		for y, val := range arr {
			if val == '0' {
				continue
			}
			currPoint := point{x, y}
			if y > 0 && arr[y-1] != '0' {
				adjacentPoint := point{x, y - 1}
				pointToIsland[currPoint] = pointToIsland[adjacentPoint]
				continue
			}
			islandId++
			pointToIsland[currPoint] = islandId
			if x > 0 && grid[x-1][y] != '0' {
				adjacentPoint := point{x - 1, y}
				if pointToIsland[adjacentPoint] > pointToIsland[currPoint] {
					pointToIsland[adjacentPoint] = pointToIsland[currPoint]
				} else if pointToIsland[adjacentPoint] < pointToIsland[currPoint] {
					pointToIsland[currPoint] = pointToIsland[adjacentPoint]
				}
			}
		}
	}

	islandIds := make([]int, 0)
	for _, id := range pointToIsland {
		islandIds = append(islandIds, id)
	}

	sort.Ints(islandIds)
	counter := 0
	for i, id := range islandIds {
		if i == 0 || islandIds[i-1] != id {
			counter++
		}
	}

	return counter
}
