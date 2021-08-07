package _1932_merge_bsts_to_create_single_bst

import (
	"fmt"
	"github.com/babadro/algorithms-go/base/binaryTree"
	"reflect"
	"testing"
)

func Test_canMerge(t *testing.T) {
	tests := []struct {
		trees [][]int
		want  func() *binaryTree.Node
	}{
		{
			trees: [][]int{{2, 1}, {3, 2, 5}, {5, 4}},
			want: func() *binaryTree.Node {
				n1, n2, n3, n4, n5 := n(1), n(2), n(3), n(4), n(5)

				n3.Left, n3.Right = n2, n5
				n2.Left = n1
				n5.Left = n4

				return n3
			},
		},
		{
			trees: [][]int{{5, 3, 8}, {3, 2, 6}},
			want: func() *binaryTree.Node {
				return nil
			},
		},
		{
			trees: [][]int{{5, 4}, {3}},
			want: func() *binaryTree.Node {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.trees), func(t *testing.T) {
			trees := arrsToTrees(tt.trees)
			wantTree := tt.want()
			if got := canMerge(trees); !reflect.DeepEqual(got, wantTree) {
				t.Errorf("canMerge() = %v, want %v", got, wantTree)
			}
		})
	}
}

func arrsToTrees(vals [][]int) []*binaryTree.Node {
	res := make([]*binaryTree.Node, len(vals))
	for i, arr := range vals {
		res[i] = newTree(arr)
	}

	return res
}

func newTree(vals []int) *binaryTree.Node {
	node := &binaryTree.Node{Val: vals[0]}
	if len(vals) > 1 {
		node.Left = &binaryTree.Node{Val: vals[1]}
	}
	if len(vals) > 2 {
		node.Right = &binaryTree.Node{Val: vals[2]}
	}

	return node
}

func n(v int) *binaryTree.Node {
	return &binaryTree.Node{Val: v}
}
