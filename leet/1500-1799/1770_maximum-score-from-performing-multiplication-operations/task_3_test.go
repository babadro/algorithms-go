package _1770_maximum_score_from_performing_multiplication_operations

import "testing"

func Test_maximumScore2(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		multipliers []int
		want        int
	}{
		{"1", []int{1, 2, 3}, []int{3, 2, 1}, 14},
		{
			"2",
			[]int{-5, -3, -3, -2, 7, 1},
			[]int{-10, -5, 3, 4, 6},
			102,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.nums, tt.multipliers); got != tt.want {
				t.Errorf("maximumScore2() = %v, want %v", got, tt.want)
			}
		})
	}
}
