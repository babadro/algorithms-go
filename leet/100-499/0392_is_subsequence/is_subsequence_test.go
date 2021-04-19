package _392_is_subsequence

import "testing"

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"1", "abc", "ahbgdc", true},
		{"2", "axc", "ahbgdc", false},
		{"3", "abc", "aba", false},
		{"4", "", "aba", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.s, tt.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
