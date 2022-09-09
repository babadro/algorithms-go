package __level_order_successor

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
	"github.com/stretchr/testify/require"
)

func Test_findSuccessor(t *testing.T) {
	tests := []struct {
		buildTree func() *binaryTree.Node
		key       int
		want      *int
	}{
		{
			buildTree: func() *binaryTree.Node {
				n1, n2, n3, n4, n5 := n(1), n(2), n(3), n(4), n(5)
				n1.Left, n1.Right = n2, n3
				n2.Left, n2.Right = n4, n5

				return n1
			}, key: 3, want: ptr(4),
		},
		{
			buildTree: func() *binaryTree.Node {
				n1, n5, n7, n9, n10, n12 := n(1), n(5), n(7), n(9), n(10), n(12)
				n12.Left, n12.Right = n7, n1
				n7.Left = n9
				n1.Left, n1.Right = n10, n5

				return n12
			}, key: 9, want: ptr(10),
		},
		{
			buildTree: func() *binaryTree.Node {
				n1, n5, n7, n9, n10, n12 := n(1), n(5), n(7), n(9), n(10), n(12)
				n12.Left, n12.Right = n7, n1
				n7.Left = n9
				n1.Left, n1.Right = n10, n5

				return n12
			}, key: 12, want: ptr(7),
		},
		{
			buildTree: func() *binaryTree.Node {
				return n(1)
			}, key: 1, want: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := findSuccessor(tt.buildTree(), tt.key)
			if tt.want == nil {
				require.Nil(t, got)
			} else {
				require.Equal(t, *tt.want, got.Val)
			}
		})
	}
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}

func ptr(val int) *int {
	return &val
}
