package lexicographic_rank_of_string_with_duplicates

import "testing"

func Test_findRank(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abab", 2},
		{"settLe", 107},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := findRank(tt.s); got != tt.want {
				t.Errorf("findRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
