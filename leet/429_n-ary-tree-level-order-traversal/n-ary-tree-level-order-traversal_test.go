package _429_n_ary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name        string
		treeBuilder func() *Node
		want        [][]int
	}{
		{
			"1",
			func() *Node {
				n1, n2, n3, n4, n5, n6 := n(1), n(2), n(3), n(4), n(5), n(6)

				n1.Children = []*Node{n3, n2, n4}
				n3.Children = []*Node{n5, n6}

				return n1
			},
			[][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			"2",
			func() *Node {
				return n(1)
			},
			[][]int{{1}},
		},
		{
			"3",
			func() *Node {
				n1 := n(1)
				n2, n3, n4, n5 := n(2), n(3), n(4), n(5)
				n6, n7, n8, n9, n10 := n(6), n(7), n(8), n(9), n(10)
				n11, n12, n13 := n(11), n(12), n(13)
				n14 := n(14)

				n1.Children = []*Node{n2, n3, n4, n5}
				n3.Children = []*Node{n6, n7}
				n4.Children = []*Node{n8}
				n5.Children = []*Node{n9, n10}
				n7.Children = []*Node{n11}
				n8.Children = []*Node{n12}
				n9.Children = []*Node{n13}
				n11.Children = []*Node{n14}

				return n1
			},
			[][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
		{
			"4",
			func() *Node {
				n1, n2, n3 := n(1), n(2), n(3)

				n1.Children = []*Node{n2}
				n2.Children = []*Node{n3}

				return n1
			},
			[][]int{{1}, {2}, {3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.treeBuilder()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *Node {
	return &Node{Val: val}
}
