package _058_length_of_last_word

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", " dfs a fdaf   ", 4},
		{"2", "    ", 0},
		{"3", "fff", 3},
		{"4", " fff ", 3},
		{"5", " fff", 3},
		{"6", "fff ", 3},
		{"7", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
