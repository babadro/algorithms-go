package _036_valid_sudoku

// TODO 1
func isValidSudoku(board [][]byte) bool {
	arr := [9]bool{}
	for _, row := range board {
		for k := range arr {
			arr[k] = false
		}
		for _, v := range row {
			if !valid(arr, v) {
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		for k := range arr {
			arr[k] = false
		}
		for j := 0; j < 9; j++ {
			if !valid(arr, board[i][j]) {
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		for k := range arr {
			arr[k] = false
		}
		for j := 0; j < 0; j++ {
			x, y := (i*3)%3+j%3, i*3+j/3
			if !valid(arr, board[y][x]) {
				return false
			}
		}
	}
	return true
}

func valid(arr [9]bool, v byte) bool {
	if v != '.' {
		if idx := v - '0'; idx < 0 || idx > 8 || arr[idx] == true {
			return false
		} else {
			arr[idx] = true
		}
	}
	return true
}
