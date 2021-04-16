package _921_minimum_add_to_make_parentheses_valid

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	tests := []struct {
		S    string
		want int
	}{
		{"())", 1},
		{"(((", 3},
		{"()", 0},
		{"()))((", 4},
	}
	for _, tt := range tests {
		t.Run(tt.S, func(t *testing.T) {
			if got := minAddToMakeValid(tt.S); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
