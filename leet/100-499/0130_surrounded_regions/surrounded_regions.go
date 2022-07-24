package _130_surrounded_regions

// 96% 83% tptl. medium (hard for me)
func solve(board [][]byte) {
	height := len(board)
	if height == 0 {
		return
	}
	width := len(board[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == 0 || y == height-1 || x == 0 || x == width-1 {
				crawl(board, x, y)
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] == 'O' {
				board[y][x] = 'X'
			} else if board[y][x] == 'N' {
				board[y][x] = 'O'
			}
		}
	}
}

func crawl(board [][]byte, x, y int) {
	if x >= len(board[0]) || x < 0 || y >= len(board) || y < 0 || board[y][x] != 'O' {
		return
	}

	board[y][x] = 'N'

	crawl(board, x-1, y)
	crawl(board, x+1, y)
	crawl(board, x, y-1)
	crawl(board, x, y+1)
}

// Doesn't work
func solve1(board [][]byte) {
	height := len(board)
	if height == 0 {
		return
	}
	width := len(board[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] == 'O' {
				if (y == 0 || y == height-1 || x == 0 || x == width-1) ||
					board[y][x-1] == 'N' ||
					board[y-1][x] == 'N' {
					board[y][x] = 'N'
				}
			}
		}
	}

	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			if board[y][x] == 'O' && (board[y+1][x] == 'N' || board[y][x+1] == 'N') {
				board[y][x] = 'N'
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] == 'O' && (board[y+1][x] == 'N' || board[y][x+1] == 'N') {
				board[y][x] = 'N'
			}
		}
	}

	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			if board[y][x] == 'O' && (board[y-1][x] == 'N' || board[y][x-1] == 'N') {
				board[y][x] = 'N'
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] == 'N' {
				board[y][x] = 'O'
			} else if board[y][x] == 'O' {
				board[y][x] = 'X'
			}
		}
	}
}
