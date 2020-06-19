package _079_word_search

import "testing"

func TestExist(t *testing.T) {
	board1 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	board2 := [][]byte{
		{'h', 'c', 'b', 'c'},
		{'c', 'b', 'a', 'd'},
		{'o', 'c', 'b', 'd'},
		{'e', 'a', 'c', 'e'},
	}

	board3 := [][]byte{
		{'a', 'b', 'c', 'd'},
	}

	board4 := [][]byte{
		{'a'},
		{'b'},
		{'c'},
		{'d'},
	}

	board5 := [][]byte{
		{'a'},
	}

	board6 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	board7 := [][]byte{
		{'A', 'B', 'C', 'D'},
		{'E', 'F', 'G', 'H'},
		{'I', 'J', 'K', 'L'},
	}

	cases := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{board1, "ABCB", false},
		{board1, "SEE", true},
		{board1, "ABCCED", true},
		{board2, "abcd", true},
		{board2, "abdb", false},
		{board3, "abcd", true},
		{board3, "dcba", true},
		{board4, "dcba", true},
		{board4, "abcd", true},
		{board4, "aba", false},
		{board5, "aba", false},
		{board5, "b", false},
		{board5, "a", true},
		{board6, "ABCESEEEFS", true},
		{board7, "EABFGCDHLK", true},
	}

	for i, c := range cases {
		if fact := exist(c.board, c.word); fact != c.expected {
			t.Errorf("case#%d want %t, got %t for word %s", i, c.expected, fact, c.word)
		}
	}
}
