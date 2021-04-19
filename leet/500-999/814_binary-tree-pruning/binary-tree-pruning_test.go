package _814_binary_tree_pruning

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"reflect"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	tests := []struct {
		name      string
		inputTree func() *binaryTree.Node
		wantTree  func() *binaryTree.Node
	}{
		{
			name: "1",
			inputTree: func() *binaryTree.Node {
				n1_1 := n(1)
				n2_1 := n(0)
				n3_1, n3_2 := n(0), n(1)

				n1_1.Right = n2_1
				n2_1.Left, n2_1.Right = n3_1, n3_2

				return n1_1
			},
			wantTree: func() *binaryTree.Node {
				n1_1, n2_1, n3_1 := n(1), n(0), n(1)
				n1_1.Right = n2_1
				n2_1.Right = n3_1

				return n1_1
			},
		},
		{
			name: "2",
			inputTree: func() *binaryTree.Node {
				l1 := nodes(1)
				l2 := nodes(0, 1)
				l3 := nodes(0, 0, 0, 1)

				l1[0].Left, l1[0].Right = l2[0], l2[1]
				l2[0].Left, l2[0].Right = l3[0], l3[1]
				l2[1].Left, l2[1].Right = l3[2], l3[3]

				return l1[0]
			},
			wantTree: func() *binaryTree.Node {
				n1, n2, n3 := n(1), n(1), n(1)
				n2.Right = n3
				n1.Right = n2

				return n1
			},
		},
		{
			name: "3",
			inputTree: func() *binaryTree.Node {
				l1 := nodes(1)
				l2 := nodes(1, 0)
				l3 := nodes(1, 1, 0, 1)
				l4 := nodes(0)

				l1[0].Left, l1[0].Right = l2[0], l2[1]
				l2[0].Left, l2[0].Right = l3[0], l3[1]

				l2[1].Left, l2[1].Right = l3[2], l3[3]

				l3[0].Left = l4[0]

				return l1[0]
			},
			wantTree: func() *binaryTree.Node {
				l1 := nodes(1)
				l2 := nodes(1, 0)
				l3 := nodes(1, 1, 1)

				l1[0].Left, l1[0].Right = l2[0], l2[1]

				l2[0].Left, l2[0].Right = l3[0], l3[1]

				l2[1].Right = l3[2]

				return l1[0]
			},
		},
		{
			name: "4",
			inputTree: func() *binaryTree.Node {
				n1 := n(1)
				n2 := n(0)

				n1.Right = n2

				return n1
			},
			wantTree: func() *binaryTree.Node {
				return n(1)
			},
		},
		{
			name: "5",
			inputTree: func() *binaryTree.Node {
				n1, n2 := n(0), n(1)
				n1.Left = n2

				return n1
			},
			wantTree: func() *binaryTree.Node {
				n1, n2 := n(0), n(1)
				n1.Left = n2

				return n1
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.wantTree()
			var gotList, wantList []int

			got := pruneTree(tt.inputTree())
			binaryTree.InOrder(got, &gotList)
			binaryTree.InOrder(want, &wantList)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("pruneTree() = %v, want %v", gotList, wantList)
			}
		})
	}
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}

func nodes(values ...int) []*binaryTree.Node {
	res := make([]*binaryTree.Node, len(values))
	for i, val := range values {
		res[i] = &binaryTree.Node{Val: val}
	}

	return res
}
