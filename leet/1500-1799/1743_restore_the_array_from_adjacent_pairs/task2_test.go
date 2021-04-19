package _1743_restore_the_array_from_adjacent_pairs

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_restoreArray(t *testing.T) {
	tests := []struct {
		adjacentPairs [][]int
		want          []int
	}{
		//{
		//	[][]int{{2, 1}, {3, 4}, {3, 2}},
		//	[]int{4, 3, 2, 1},
		//},
		{
			[][]int{{4, -2}, {1, 4}, {-3, 1}},
			[]int{-2, 4, 1, -3},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.adjacentPairs), func(t *testing.T) {
			if got := restoreArray(tt.adjacentPairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
