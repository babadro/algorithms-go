package _334_increasing_triplet_subsequence

import "testing"

func Test_increasingTriplet(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"1", []int{1, 2, 3, 4, 5}, true},
		{"2", []int{5, 4, 3, 2, 1}, false},
		{"3", []int{}, false},
		{"4", []int{1}, false},
		{"5", []int{1, 2}, false},
		{"6", []int{1, 2, 3}, true},
		{"7", []int{1, 2, 2}, false},
		{"8", []int{1, 2, 2, 4, 5}, true},
		{"9", []int{5, 1, 5, 5, 2, 5, 4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
