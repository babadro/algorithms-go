package _658_find_k_closest_elements

import (
	"reflect"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{0, 2}, 1, 1},
		{[]int{0}, 1, 0},
		{[]int{1}, 2, 0},
		{[]int{0, 1, 2, 3}, 4, 3},
		{[]int{0, 1, 3, 4}, 2, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := binarySearch(tt.arr, tt.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findClosestElements(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		x    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
		{[]int{1, 3, 5, 7, 9}, 3, 4, []int{1, 3, 5}},
		{[]int{1}, 1, 4, []int{1}},
		{[]int{1}, 1, 0, []int{1}},

		{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5, []int{3, 3, 4}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findClosestElements(tt.arr, tt.k, tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
