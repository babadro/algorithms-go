package _110_balanced_binary_tree

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		name        string
		treeBuilder func() *btree.Node
		want        bool
	}{
		{
			name: "1",
			treeBuilder: func() *btree.Node {
				return nil
			},
			want: true,
		},
		{
			name: "2",
			treeBuilder: func() *btree.Node {
				return &btree.Node{Val: 0}
			},
			want: true,
		},
		{
			name: "3",
			treeBuilder: func() *btree.Node {
				node3 := &btree.Node{Val: 3}
				node9 := &btree.Node{Val: 9}
				node20 := &btree.Node{Val: 20}
				node15 := &btree.Node{Val: 15}
				node7 := &btree.Node{Val: 7}

				node3.Left, node3.Right = node9, node20
				node20.Left, node20.Right = node15, node7

				return node3
			},
			want: true,
		},
		{
			name: "4",
			treeBuilder: func() *btree.Node {
				n1 := &btree.Node{Val: 1}
				n2 := &btree.Node{Val: 2}
				n2_ := &btree.Node{Val: 2}
				n3 := &btree.Node{Val: 3}
				n3_ := &btree.Node{Val: 3}
				n4 := &btree.Node{Val: 4}
				n4_ := &btree.Node{Val: 4}

				n1.Left, n1.Right = n2, n2_
				n2.Left, n2.Right = n3, n3_
				n3.Left, n3.Right = n4, n4_

				return n1
			},
			want: false,
		},
		{
			name: "5",
			treeBuilder: func() *btree.Node {
				n1 := &btree.Node{Val: 1}
				n2 := &btree.Node{Val: 2}
				n3 := &btree.Node{Val: 3}

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
