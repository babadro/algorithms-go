package _993_cousins_in_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"testing"
)

func Test_isCousins(t *testing.T) {
	tests := []struct {
		name        string
		treeBuilder func() *binaryTree.Node
		x           int
		y           int
		want        bool
	}{
		{
			name: "1",
			treeBuilder: func() *binaryTree.Node {
				n1, n2, n3, n4 := n(1), n(2), n(3), n(4)
				n1.Left, n1.Right = n2, n3
				n2.Left = n4

				return n1
			},
			x:    4,
			y:    3,
			want: false,
		},
		{
			name: "2",
			treeBuilder: func() *binaryTree.Node {
				n1, n2, n3, n4, n5 := n(1), n(2), n(3), n(4), n(5)
				n1.Left, n1.Right = n2, n3
				n2.Right = n4
				n3.Right = n5

				return n1
			},
			x:    5,
			y:    4,
			want: true,
		},
		{
			name: "3",
			treeBuilder: func() *binaryTree.Node {
				n1, n2, n3, n4 := n(1), n(2), n(3), n(4)
				n1.Left, n1.Right = n2, n3
				n2.Right = n4

				return n1
			},
			x:    2,
			y:    3,
			want: false,
		},
		{
			name: "two nodes",
			treeBuilder: func() *binaryTree.Node {
				n1, n2 := n(1), n(2)
				n1.Left = n2

				return n1
			},
			x:    1,
			y:    2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins2(tt.treeBuilder(), tt.x, tt.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}
