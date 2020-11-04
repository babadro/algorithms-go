package _665_non_decreasing_array

import "testing"

func Test_checkPossibility(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"1", []int{4, 2, 3}, true},
		{"2", []int{4, 2, 1}, false},
		{"3", []int{3, 4, 2, 3}, false},
		{"4", []int{2, 4, 2, 2}, true},
		{"5", []int{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
