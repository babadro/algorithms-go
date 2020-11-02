package _221_split_string_in_balanced_strings

import "testing"

func Test_balancedStringSplit(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", "RLRRLLRLRL", 4},
		{"2", "RLLLLRRRLR", 3},
		{"3", "LLLLRRRR", 1},
		{"4", "RLRRRLLRLL", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedStringSplit(tt.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
