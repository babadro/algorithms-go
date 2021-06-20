package _905_count_sub_islands

// passed. tptl. #array
// todo 2 find shorter solution (maybe and faster too)
// https://leetcode.com/problems/count-sub-islands/discuss/1284232/Fill-Second-and-Count-First
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	counter := 2
	for y := 0; y < len(grid1); y++ {
		for x := 0; x < len(grid1[0]); x++ {
			if helper1(y, x, grid1, counter) {
				counter++
			}
		}
	}

	res := 0
	for y := 0; y < len(grid2); y++ {
		for x := 0; x < len(grid2[0]); x++ {
			if grid2[y][x] != 1 {
				continue
			}

			id := 0
			helper2(y, x, grid1, grid2, &id)

			if id != -1 {
				res++
			}
		}
	}

	return res
}

func helper1(y, x int, grid [][]int, id int) bool {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
		return false
	}

	if grid[y][x] != 1 {
		return false
	}

	grid[y][x] = id

	helper1(y-1, x, grid, id)
	helper1(y+1, x, grid, id)
	helper1(y, x-1, grid, id)
	helper1(y, x+1, grid, id)

	return true
}

func helper2(y, x int, grid1, grid2 [][]int, id *int) {
	if y < 0 || y >= len(grid2) || x < 0 || x >= len(grid2[0]) {
		return
	}

	if grid2[y][x] != 1 {
		return
	}

	grid2[y][x] = 2

	grid1ID := grid1[y][x]

	if grid1ID == 0 || (*id > 0 && grid1ID != *id) {
		*id = -1
	} else if *id == 0 {
		*id = grid1ID
	}

	helper2(y-1, x, grid1, grid2, id)
	helper2(y+1, x, grid1, grid2, id)
	helper2(y, x-1, grid1, grid2, id)
	helper2(y, x+1, grid1, grid2, id)

	return
}
