package _2265_count_nodes_equal_to_average_of_subtree

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_averageOfSubtree(t *testing.T) {
	tests := []struct {
		rootBuilder func() *binaryTree.Node
		want        int
	}{
		{
			rootBuilder: func() *binaryTree.Node {
				n4, n8, n5, n0, n1, n6 := n(4), n(8), n(5), n(0), n(1), n(6)

				n4.Left, n4.Right = n8, n5

				n8.Left, n8.Right = n0, n1
				n5.Right = n6

				return n4
			},
			want: 5,
		},
		{
			rootBuilder: func() *binaryTree.Node {
				return n(1)
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := averageOfSubtree(tt.rootBuilder()); got != tt.want {
				t.Errorf("averageOfSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(n int) *binaryTree.Node {
	return &binaryTree.Node{Val: n}
}
