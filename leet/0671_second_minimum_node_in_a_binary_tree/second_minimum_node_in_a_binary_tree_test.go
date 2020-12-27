package _671_second_minimum_node_in_a_binary_tree

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func Test_findSecondMinimumValue(t *testing.T) {
	tests := []struct {
		name        string
		treeBuilder func() *btree.Node
		want        int
	}{
		//{
		//	name: "1",
		//	treeBuilder: func() *btree.Node {
		//		n1_1 := n(1)
		//		n1_2, n3_2 := n(1), n(3)
		//		n1_3_1, n1_3_2, n3_3, n4_3 := n(1), n(1), n(3), n(4)
		//		n3_4, n1_4_1, n1_4_2, n1_4_3 := n(3), n(1), n(1), n(1)
		//		n3_5_1, n3_5_2, n1_5_1, n6_5, n2_5, n1_5_2 := n(3), n(3), n(1), n(6), n(2), n(1)
		//
		//		n1_1.Left, n1_1.Right = n1_2, n3_2
		//		n1_2.Left, n1_2.Right = n1_3_1, n1_3_2
		//		n1_3_1.Left, n1_3_1.Right = n3_4, n1_4_1
		//	},
		//	want: -1,
		//},
		{
			name: "2",
			treeBuilder: func() *btree.Node {
				return &btree.Node{Val: 0}
			},
			want: -1,
		},
		{
			name: "3",
			treeBuilder: func() *btree.Node {
				node1 := &btree.Node{Val: 2}
				node2 := &btree.Node{Val: 2}
				node3 := &btree.Node{Val: 2}

				node1.Left, node1.Right = node2, node3

				return node1
			},
			want: -1,
		},
		{
			name: "4",
			treeBuilder: func() *btree.Node {
				n2 := &btree.Node{Val: 2}
				n2_ := &btree.Node{Val: 2}
				n5 := &btree.Node{Val: 5}
				n5_ := &btree.Node{Val: 5}
				n7 := &btree.Node{Val: 7}

				n2.Left, n2.Right = n2_, n5
				n5.Left, n5.Right = n5_, n7

				return n2
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue2(tt.treeBuilder()); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *btree.Node {
	return &btree.Node{Val: val}
}
