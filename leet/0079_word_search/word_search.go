package _079_word_search

// 96% 84%
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	totalCount := m * n
	if len(word) > totalCount {
		return false
	}
	for y := range board {
		for x := range board[y] {
			if success := search(y, x, m, n, 0, word, board); success {
				return true
			}
		}
	}
	return false
}

func search(y, x, m, n, idx int, word string, board [][]byte) bool {
	if y < 0 || y >= m || x < 0 || x >= n {
		return false
	}
	char := word[idx]
	if board[y][x] != char {
		return false
	}
	idx++
	if idx == len(word) {
		return true
	}
	board[y][x] = 0
	result := search(y-1, x, m, n, idx, word, board) ||
		search(y+1, x, m, n, idx, word, board) ||
		search(y, x-1, m, n, idx, word, board) ||
		search(y, x+1, m, n, idx, word, board)

	if !result {
		board[y][x] = char
	}
	return result
}
