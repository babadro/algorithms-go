package _127_word_ladder

import "testing"

func TestLadderLength(t *testing.T) {
	cases := []struct {
		beginWord, endWord string
		wordList           []string
		expected           int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
	}

	for i, c := range cases {
		if fact := ladderLength(c.beginWord, c.endWord, c.wordList); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
