package _151_reverse_words_in_a_string

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Alice does not even like bob", "bob like even not does Alice"},
		{"a", "a"},
		{"  a   ", "a"},
		{"a       b", "b a"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
