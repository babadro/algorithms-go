package _1737_change_minimum_characters_to_satisfy_one_of_three_conditions

import (
	"testing"
)

func Test_minCharacters(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"aba", "caa", 2},
		{"dabadd", "cda", 3},
		{"aaaabc", "aaad", 3},
		{"aaaabc", "az", 1},
		{"a", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", 2},
	}
	for _, tt := range tests {
		t.Run(tt.a+"_"+tt.b, func(t *testing.T) {
			if got := minCharacters(tt.a, tt.b); got != tt.want {
				t.Errorf("minCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
