package _124_binary_tree_maximum_path_sum

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *binaryTree.Node
	}

	n1, n2, n3 := n(1), n(2), n(3)
	n1.Left, n1.Right = n2, n3
	args1 := args{root: n1}

	n1, n2, n3, n4, n5 := n(-10), n(9), n(20), n(15), n(7)
	n1.Left, n1.Right = n2, n3
	n3.Left, n3.Right = n4, n5
	args2 := args{root: n1}

	n1 = n(-3)
	args3 := args{root: n1}

	tests := []struct {
		args args
		want int
	}{
		{args: args1, want: 6},
		{args2, 42},
		{args3, -3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}
