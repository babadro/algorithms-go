package __cyclic_sort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sort(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{[]int{3, 2, 1, 4, 5}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			cyclicSort(tt.nums)

			require.True(t, sort.IntsAreSorted(tt.nums))
		})
	}
}
