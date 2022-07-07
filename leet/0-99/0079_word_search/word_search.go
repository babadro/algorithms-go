package _079_word_search

// passed
func exist(board [][]byte, word string) bool {
	for y := range board {
		for x := range board[y] {
			if rec(board, y, x, 0, word) {
				return true
			}
		}
	}

	return false
}

func rec(board [][]byte, y, x, i int, word string) bool {
	if i == len(word) {
		return true
	}

	if y == len(board) || y < 0 || x == len(board[0]) || x < 0 {
		return false
	}

	char := board[y][x]
	if word[i] != char {
		return false
	}

	board[y][x] = '0'

	i++

	res := rec(board, y+1, x, i, word) ||
		rec(board, y, x+1, i, word) ||
		rec(board, y-1, x, i, word) ||
		rec(board, y, x-1, i, word)

	board[y][x] = char

	return res
}
