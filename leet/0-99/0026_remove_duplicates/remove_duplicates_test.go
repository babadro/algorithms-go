package _026_remove_duplicates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates3(t *testing.T) {
	tests := []struct {
		nums     []int
		wantNums []int
		wantK    int
	}{
		{[]int{1, 2, 2, 3}, []int{1, 2, 3, 3}, 3},
		{[]int{1, 1}, []int{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			res := removeDuplicates(tt.nums)
			require.Equal(t, tt.wantK, res)
			require.Equal(t, tt.wantNums, tt.nums)
		})
	}
}
