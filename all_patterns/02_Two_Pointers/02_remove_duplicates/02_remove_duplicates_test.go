package _2_remove_duplicates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		sortedArr []int
		wantArr   []int
	}{
		{[]int{2, 3, 3, 3, 6, 9, 9}, []int{2, 3, 6, 9}},
		{[]int{2, 2, 2, 11}, []int{2, 11}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.sortedArr), func(t *testing.T) {
			end := removeDuplicates(tt.sortedArr)
			require.Equal(t, tt.wantArr, tt.sortedArr[:end])
		})
	}
}

func Test_remove(t *testing.T) {
	tests := []struct {
		arr     []int
		key     int
		wantArr []int
	}{
		{[]int{3, 2, 3, 6, 3, 10, 9, 3}, 3, []int{2, 6, 10, 9}},
		{[]int{2, 11, 2, 2, 1}, 2, []int{11, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %d", tt.arr, tt.key), func(t *testing.T) {
			end := remove(tt.arr, tt.key)
			require.Equal(t, tt.wantArr, tt.arr[:end])
		})
	}
}
