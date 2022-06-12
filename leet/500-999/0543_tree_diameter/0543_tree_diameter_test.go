package _543_tree_diameter

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *binaryTree.Node
	}

	n1, n2, n3, n4, n5 := n(1), n(2), n(3), n(4), n(5)
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	args1 := args{root: n1}

	n1, n2 = n(1), n(2)
	n1.Left = n2
	args2 := args{root: n1}

	tests := []struct {
		args args
		want int
	}{
		{args: args1, want: 3},
		{args2, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}
