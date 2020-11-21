package _720_longest_word_in_dictionary

import "testing"

func Test_longestWord(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{"1", []string{"w", "wo", "wor", "worl", "world"}, "world"},
		{"2", []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"},
		{"3", []string{"a", "b", "c"}, "a"},
		{"4", []string{"a", "ab", "c"}, "ab"},
		{"5", []string{"ts", "e", "x", "pbhj", "opto", "xhigy", "erikz", "pbh", "opt", "erikzb", "eri", "erik", "xlye", "xhig", "optoj", "optoje", "xly", "pb", "xhi", "x", "o"}, "e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWord(tt.words); got != tt.want {
				t.Errorf("longestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
