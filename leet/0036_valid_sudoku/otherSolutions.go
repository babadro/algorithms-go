package _036_valid_sudoku

func isValidSudoku2(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' && isOK(i, j, board) == false {
				return false
			}
		}
	}

	return true
}

func isOK(i, j int, board [][]byte) bool {
	row := 3 * (i / 3)
	col := 3 * (j / 3)
	c := board[i][j]
	board[i][j] = '.' // avoid compare with self
	defer func() {
		board[i][j] = c // restore
	}()
	for k := 0; k < 9; k++ {
		if board[k][j] == c { // check row
			return false
		}
		if board[i][k] == c { // check col
			return false
		}
		if board[row+k/3][col+k%3] == c { // check block
			return false
		}
	}
	return true
}

func isValidSudoku3(board [][]byte) bool {
	if len(board) == 0 || board == nil {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' && !isValidSudokuOne(board, i, j, board[i][j]) {
				return false
			}
		}
	}
	return true
}

func isValidSudokuOne(board [][]byte, row, col int, val byte) bool {
	for i := 0; i < 9; i++ {
		if i != row && board[i][col] != '.' && board[i][col] == val {
			return false
		}
		if i != col && board[row][i] != '.' && board[row][i] == val {
			return false
		}
		x := 3*(row/3) + i/3
		y := 3*(col/3) + i%3
		if x == row && y == col {
			continue
		}
		if board[x][y] != '.' && board[x][y] == val {
			return false
		}
	}
	return true
}

func isValidSudoku4(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' && isOK2(i, j, board) == false {
				return false
			}
		}
	}

	return true
}

func isOK2(i, j int, board [][]byte) bool {
	row := 3 * (i / 3)
	col := 3 * (j / 3)
	c := board[i][j]
	board[i][j] = '.' // avoid compare with self
	defer func() {
		board[i][j] = c // restore
	}()
	for k := 0; k < 9; k++ {
		if board[k][j] == c { // check row
			return false
		}
		if board[i][k] == c { // check col
			return false
		}
		if board[row+k/3][col+k%3] == c { // check block
			return false
		}
	}
	return true
}
