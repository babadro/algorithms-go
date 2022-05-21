package _1095_find_in_mountain_array

import (
	"fmt"
	"testing"
)

func Test_findInMountainArray(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5, 3, 1}, 3, 2},
		{[]int{0, 1, 2, 4, 2, 1}, 3, -1},
		{[]int{1, 2, 1}, 1, 0},
		{[]int{1, 5, 2}, 2, 2},
		{bigInput1, 405001, 4050},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			m := MountainArray{arr: tt.arr}
			if got := findInMountainArray(tt.target, &m); got != tt.want {
				t.Errorf("findInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
