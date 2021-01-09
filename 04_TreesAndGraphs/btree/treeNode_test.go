package btree

import (
	"reflect"
	"testing"
)

func TestLevelOrderFunc(t *testing.T) {
	tests := []struct {
		name        string
		treeBuilder func() *Node
		want        []int
	}{
		{
			name: "1",
			treeBuilder: func() *Node {
				n1, n2, n3, n4, n5 := n(1), n(2), n(3), n(4), n(5)
				n1.Left, n1.Right = n2, n3
				n2.Left, n2.Right = n4, n5

				return n1
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "single node",
			treeBuilder: func() *Node {
				return n(1)
			},
			want: []int{1},
		},
		{
			name: "left linked list",
			treeBuilder: func() *Node {
				n1, n2, n3, n4, n5 := n(1), n(2), n(3), n(4), n(5)
				n1.Left = n2
				n2.Left = n3
				n3.Left = n4
				n4.Left = n5

				return n1
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := make([]int, 0)
			f := func(node *Node) {
				list = append(list, node.Val)
			}

			LevelOrderFuncIterative(tt.treeBuilder(), f)

			if !reflect.DeepEqual(list, tt.want) {
				t.Errorf("got = %v, want %v", list, tt.want)
			}
		})
	}
}

func n(val int) *Node {
	return &Node{Val: val}
}
