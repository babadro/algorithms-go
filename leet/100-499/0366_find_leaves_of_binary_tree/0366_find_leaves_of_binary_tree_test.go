package _366_find_leaves_of_binary_tree

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
	"github.com/stretchr/testify/require"
)

func Test_findLeaves(t *testing.T) {
	tests := []struct {
		tree []int
		want [][]int
	}{
		{
			[]int{1, 2, 3, 4, 5}, [][]int{{4, 5, 3}, {2}, {1}},
		},
		{
			[]int{1}, [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			root := binaryTree.ArrayToBinaryTree(tt.tree)

			res := findLeaves(root)

			require.Equal(t, res, tt.want)
		})
	}
}
