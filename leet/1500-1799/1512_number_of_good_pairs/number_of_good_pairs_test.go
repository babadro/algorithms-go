package _512_number_of_good_pairs

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1", []int{1, 2, 3, 1, 1, 3}, 4},
		{"2", []int{1, 1, 1, 1}, 6},
		{"3", []int{1, 2, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIdenticalPairs(tt.nums); got != tt.want {
				t.Errorf("numIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
