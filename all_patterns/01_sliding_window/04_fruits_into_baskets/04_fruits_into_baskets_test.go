package _4_fruits_into_baskets

import "testing"

func Test_findMaxNumFruits(t *testing.T) {
	tests := []struct {
		arr  string
		want int
	}{
		{"ABCAC", 3},
		{"ABCBBC", 5},
		{"ABABABABABABCCCCCCCCCCCC", 13},
	}
	for _, tt := range tests {
		t.Run(tt.arr, func(t *testing.T) {
			if got := findMaxNumFruits(tt.arr); got != tt.want {
				t.Errorf("findMaxNumFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
