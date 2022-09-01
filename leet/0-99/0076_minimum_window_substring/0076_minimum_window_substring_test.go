package _076_minimum_window_substring

import "testing"

func Test_minWindow(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minWindow(tt.s, tt.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
