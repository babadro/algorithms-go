package _5_search_in_a_sorted_infinite_array

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	tests := []struct {
		arr  []int
		key  int
		want int
	}{
		{[]int{4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}, 16, 6},
		{[]int{4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}, 11, -1},
		{[]int{1, 3, 8, 10, 15}, 15, 4},
		{[]int{1, 3, 8, 10, 15}, 20, -1},
		{[]int{}, 2, -1},
	}
	for _, tt := range tests {
		reader := arrayReader{arr: tt.arr}
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if got := search(reader, tt.key); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
