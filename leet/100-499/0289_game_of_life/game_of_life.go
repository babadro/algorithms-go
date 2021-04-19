package _289_game_of_life

var neighborsIdxes = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

// 100% 84%
// TODO 2 read Follow up 2 : Infinite Board in the solution:
// https://leetcode.com/problems/game-of-life/solution/
func gameOfLife(board [][]int) {
	for m := range board {
		for n := range board[m] {
			aliveNeighbors := 0
			for _, pair := range neighborsIdxes {
				y, x := m+pair[0], n+pair[1]
				if y < 0 || y >= len(board) || x < 0 || x >= len(board[m]) {
					continue
				}
				if board[y][x] == 1 || board[y][x] == 2 {
					aliveNeighbors++
				}
			}

			if board[m][n] == 1 {
				if aliveNeighbors != 2 && aliveNeighbors != 3 {
					board[m][n] = 2
				}
			} else {
				if aliveNeighbors == 3 {
					board[m][n] = 3
				}
			}
		}
	}

	for m := range board {
		for n := range board[m] {
			if board[m][n] == 2 {
				board[m][n] = 0
			} else if board[m][n] == 3 {
				board[m][n] = 1
			}
		}
	}
}
