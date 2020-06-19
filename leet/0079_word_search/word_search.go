package _079_word_search

// 26%, 54%
// TODO 2 look discuss for best solution
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	totalCount := m * n
	if len(word) > totalCount {
		return false
	}
	visited := make(map[key]bool)
	for y := range board {
		for x := range board[y] {
			if success := search(y, x, m, n, 0, word, board, visited); success {
				return true
			}
		}
	}
	return false
}

func search(y, x, m, n, idx int, word string, board [][]byte, visited map[key]bool) bool {
	if y < 0 || y >= m || x < 0 || x >= n {
		return false
	}
	if board[y][x] != word[idx] {
		return false
	}
	k := key{byte(x), byte(y)}
	if visited[k] {
		return false
	}
	idx++
	if idx == len(word) {
		return true
	}
	visited[k] = true
	result := search(y-1, x, m, n, idx, word, board, visited) ||
		search(y+1, x, m, n, idx, word, board, visited) ||
		search(y, x-1, m, n, idx, word, board, visited) ||
		search(y, x+1, m, n, idx, word, board, visited)

	if !result {
		visited[k] = false
	}
	return result
}

type key struct {
	x, y byte
}
