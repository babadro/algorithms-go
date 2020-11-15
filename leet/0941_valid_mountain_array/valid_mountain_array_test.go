package _941_valid_mountain_array

import "testing"

func Test_validMountainArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{"123", []int{1, 2, 1}, true},
		{"202", []int{2, 0, 2}, false},
		{"", []int{}, false},
		{"2222", []int{2, 2, 2, 2}, false},
		{"1234543", []int{1, 2, 3, 4, 5, 4, 3}, true},
		{"12334543", []int{1, 2, 3, 3, 4, 5, 4, 3}, false},
		{"123321", []int{1, 2, 3, 3, 2, 1}, false},
		{"1234", []int{1, 2, 3, 4}, false},
		{"4321", []int{4, 3, 2, 1}, false},
		{"1234321234321", []int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
