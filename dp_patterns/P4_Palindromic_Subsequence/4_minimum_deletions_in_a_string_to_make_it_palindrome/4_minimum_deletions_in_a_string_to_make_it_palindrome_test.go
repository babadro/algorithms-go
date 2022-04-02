package __minimum_deletions_in_a_string_to_make_it_palindrome

import "testing"

func Test_findMinimumDeletions(t *testing.T) {
	tests := []struct {
		st   string
		want int
	}{
		{"abdbca", 1},
		{"cddpd", 2},
		{"pqr", 2},
	}
	for _, tt := range tests {
		t.Run(tt.st, func(t *testing.T) {
			if got := findMinimumDeletions(tt.st); got != tt.want {
				t.Errorf("findMinimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
