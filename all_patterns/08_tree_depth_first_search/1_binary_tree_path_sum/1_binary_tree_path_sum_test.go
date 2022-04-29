package __binary_tree_path_sum

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_hasPath(t *testing.T) {
	n1_1, n1_2, n1_3, n1_4, n1_5, n1_6, n1_7 := n(1), n(2), n(3), n(4), n(5), n(6), n(7)
	n1_1.Left, n1_1.Right = n1_2, n1_3
	n1_2.Left, n1_2.Right = n1_4, n1_5
	n1_3.Left, n1_3.Right = n1_6, n1_7

	n2_12, n2_7, n2_9, n2_1, n2_10, n2_5 := n(12), n(7), n(9), n(1), n(10), n(5)
	n2_12.Left, n2_12.Right = n2_7, n2_1
	n2_7.Left = n2_9
	n2_1.Left, n2_1.Right = n2_10, n2_5

	tests := []struct {
		root *binaryTree.Node
		sum  int
		want bool
	}{
		{n1_1, 10, true},
		{n2_12, 23, true},
		{n2_12, 16, false},
		{nil, 0, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := hasPath(tt.root, tt.sum); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}
