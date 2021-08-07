package _1932_merge_bsts_to_create_single_bst

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"reflect"
	"testing"
)

func Test_canMerge2(t *testing.T) {
	tests := []struct {
		name  string
		trees func() []*binaryTree.Node
		want  func() *binaryTree.Node
	}{
		{
			trees: func() []*binaryTree.Node {
				res := make([]*binaryTree.Node, 3)

				n2, n1 := n(2), n(1)
				n2.Left = n1
				res[0] = n2

				n2, n3, n5 := n(2), n(3), n(5)
				n3.Left, n3.Right = n2, n5
				res[1] = n3

				n5, n4 := n(5), n(4)
				n5.Left = n4
				res[2] = n5

				return res
			},
			want: func() *binaryTree.Node {
				n := nodes(1, 2, 3, 4, 5)
				n[3].Left, n[3].Right = n[2], n[5]
				n[2].Left = n[1]
				n[5].Left = n[4]

				return n[3]
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trees, want := tt.trees(), tt.want()
			if got := canMerge2(trees); !reflect.DeepEqual(got, want) {
				t.Errorf("canMerge() = %v, want %v", got, want)
			}
		})
	}
}

func nodes(values ...int) map[int]*binaryTree.Node {
	res := make(map[int]*binaryTree.Node)

	for _, v := range values {
		res[v] = n(v)
	}

	return res
}
