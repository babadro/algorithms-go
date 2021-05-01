package _1161_maximum_level_sum_of_a_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"testing"
)

func Test_maxLevelSum(t *testing.T) {
	tests := []struct {
		name        string
		treeBuilder func() *binaryTree.Node
		want        int
	}{
		//{
		//	treeBuilder: func() *binaryTree.Node {
		//		n1_1 := n(1)
		//		n2_7, n2_0 := n(7), n(0)
		//		n3_7, n3_m8 := n(7), n(-8)
		//
		//		n1_1.Left, n1_1.Right = n2_7, n2_0
		//		n2_7.Left, n2_7.Right = n3_7, n3_m8
		//
		//		return n1_1
		//	},
		//	want: 2,
		//},
		//{
		//	treeBuilder: func() *binaryTree.Node {
		//		n1, n2, n3, n4, n5, n6, n7 := n(1), n(2), n(3), n(4), n(5), n(6), n(7)
		//		n1.Left, n1.Right = n2, n3
		//		n2.Left, n2.Right = n4, n5
		//		n3.Left, n3.Right = n6, n7
		//
		//		return n1
		//	},
		//	want: 3,
		//},
		//{
		//	treeBuilder: func() *binaryTree.Node {
		//		n1, n2, n3 := n(1), n(1), n(0)
		//		n1.Left, n1.Right = n2, n3
		//
		//		return n1
		//	},
		//	want: 1,
		//},
		{
			treeBuilder: func() *binaryTree.Node {
				n1, n2, n3, n4, n5, n6 := n(-100), n(-200), n(-300), n(-20), n(-5), n(-10)
				n1.Left, n1.Right = n2, n3
				n2.Left, n2.Right = n4, n5
				n3.Left = n6

				return n1
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLevelSum(tt.treeBuilder()); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}
