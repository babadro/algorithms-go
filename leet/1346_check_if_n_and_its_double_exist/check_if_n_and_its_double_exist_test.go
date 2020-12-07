package _1346_check_if_n_and_its_double_exist

import "testing"

func Test_checkIfExist(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{"1", []int{10, 2, 5, 3}, true},
		{"2", []int{7, 1, 14, 11}, true},
		{"3", []int{3, 1, 7, 11}, false},
		{"4", []int{0, 0}, true},
		{"5", []int{-2, 0, 10, -19, 4, 6, -8}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist2(tt.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
