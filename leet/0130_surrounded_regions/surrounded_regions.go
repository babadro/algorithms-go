package _130_surrounded_regions

// TODO 1
func solve(board [][]byte) {
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
