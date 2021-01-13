package _110_balanced_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		name        string
		treeBuilder func() *binaryTree.Node
		want        bool
	}{
		{
			name: "1",
			treeBuilder: func() *binaryTree.Node {
				return nil
			},
			want: true,
		},
		{
			name: "2",
			treeBuilder: func() *binaryTree.Node {
				return &binaryTree.Node{Val: 0}
			},
			want: true,
		},
		{
			name: "3",
			treeBuilder: func() *binaryTree.Node {
				node3 := &binaryTree.Node{Val: 3}
				node9 := &binaryTree.Node{Val: 9}
				node20 := &binaryTree.Node{Val: 20}
				node15 := &binaryTree.Node{Val: 15}
				node7 := &binaryTree.Node{Val: 7}

				node3.Left, node3.Right = node9, node20
				node20.Left, node20.Right = node15, node7

				return node3
			},
			want: true,
		},
		{
			name: "4",
			treeBuilder: func() *binaryTree.Node {
				n1 := &binaryTree.Node{Val: 1}
				n2 := &binaryTree.Node{Val: 2}
				n2_ := &binaryTree.Node{Val: 2}
				n3 := &binaryTree.Node{Val: 3}
				n3_ := &binaryTree.Node{Val: 3}
				n4 := &binaryTree.Node{Val: 4}
				n4_ := &binaryTree.Node{Val: 4}

				n1.Left, n1.Right = n2, n2_
				n2.Left, n2.Right = n3, n3_
				n3.Left, n3.Right = n4, n4_

				return n1
			},
			want: false,
		},
		{
			name: "5",
			treeBuilder: func() *binaryTree.Node {
				n1 := &binaryTree.Node{Val: 1}
				n2 := &binaryTree.Node{Val: 2}
				n3 := &binaryTree.Node{Val: 3}

				n1.Left = n2
				n2.Left = n3

				return n1
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced2(tt.treeBuilder()); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
