package _036_valid_sudoku

// It's really fast
// TODO 3 можно попробовать сразу сделать 27 массивов для проверки (даже не слайсов, а массивов) и за один проход через матрицу свалидировать.
func isValidSudoku(board [][]byte) bool {
	arr := make([]bool, 9)
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
			if !valid(arr, board[j][i]) {
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		for k := range arr {
			arr[k] = false
		}
		for j := 0; j < 9; j++ {
			x, y := (i*3)%9+j%3, ((i/3)*3)+j/3
			if !valid(arr, board[y][x]) {
				return false
			}
		}
	}
	return true
}

func valid(arr []bool, v byte) bool {
	if v != '.' {
		if idx := v - '1'; idx < 0 || idx > 8 || arr[idx] == true {
			return false
		} else {
			arr[idx] = true
		}
	}
	return true
}
