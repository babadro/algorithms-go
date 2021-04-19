package _836_rectangle_overlap

import "testing"

func Test_isRectangleOverlap(t *testing.T) {
	tests := []struct {
		name string
		rec1 []int
		rec2 []int
		want bool
	}{
		{"1", []int{0, 0, 2, 2}, []int{1, 1, 3, 3}, true},
		{"2", []int{0, 0, 1, 1}, []int{1, 0, 2, 1}, false},
		{"3", []int{0, 0, 1, 1}, []int{2, 2, 3, 3}, false},
		{"4", []int{0, 0, 0, 0}, []int{0, 0, 0, 0}, false},
		{"5", []int{0, 0, 1, 1}, []int{1, 1, 2, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleOverlap(tt.rec1, tt.rec2); got != tt.want {
				t.Errorf("isRectangleOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
