package _767_reorganize_string

import "testing"

func Test_reorganizeString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"aab", "aba"},
		{"aaab", ""},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := reorganizeString(tt.s); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
