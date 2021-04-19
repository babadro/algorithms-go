package _1046_last_stone_weight

import "testing"

func Test_lastStoneWeight(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{"1", []int{2, 7, 4, 1, 8, 1}, 1},
		{"1", []int{2}, 2},
		{"1", []int{2, 3, 4, 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
