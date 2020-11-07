package _414_third_maximum_number

import "testing"

func Test_thirdMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1", []int{3, 2, 1}, 1},
		{"1", []int{1, 2}, 2},
		{"1", []int{2, 2, 3, 1}, 1},
		{"1", []int{1, 2, 3, 2, 1}, 1},
		{"1", []int{3, 3, 3, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
