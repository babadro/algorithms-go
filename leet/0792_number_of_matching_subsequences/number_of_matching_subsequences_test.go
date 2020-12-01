package _792_number_of_matching_subsequences

import "testing"

func Test_numMatchingSubseq(t *testing.T) {
	tests := []struct {
		name  string
		S     string
		words []string
		want  int
	}{
		{"1", "abcde", []string{"a", "bb", "acd", "ace"}, 3},
		{"2", "adfs", []string{"a", "df"}, 2},
		{"3", "a", []string{"a"}, 1},
		{"4", "b", []string{"a"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMatchingSubseq4(tt.S, tt.words); got != tt.want {
				t.Errorf("numMatchingSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
