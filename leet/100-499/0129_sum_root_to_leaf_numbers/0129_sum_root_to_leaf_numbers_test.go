package _129_sum_root_to_leaf_numbers

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *binaryTree.Node
	}
	n1, n2, n3 := n(1), n(2), n(3)
	n1.Left, n1.Right = n2, n3
	args1 := args{root: n1}

	n4, n9, n0, n5, n1 := n(4), n(9), n(0), n(5), n(1)
	n4.Left, n4.Right = n9, n0
	n9.Left, n9.Right = n5, n1
	args2 := args{root: n4}

	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args1, want: 25},
		{args: args2, want: 1026},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers2(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(v int) *binaryTree.Node {
	return &binaryTree.Node{Val: v}
}
