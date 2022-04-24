package _2_reverse_level_order_traversal

import (
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_reverseLevelOrderTraversal(t *testing.T) {
	n1, n2, n3, n4, n5, n6, n7 := n(1), n(2), n(3), n(4), n(5), n(6), n(7)
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Left, n3.Right = n6, n7

	n8, n9, n10, n11, n12, n13 := n(8), n(9), n(10), n(11), n(12), n(13)
	n8.Left, n8.Right = n9, n10
	n9.Right = n11
	n10.Left, n10.Right = n12, n13

	n14, n15, n16, n17 := n(14), n(15), n(16), n(17)
	n14.Left = n15
	n15.Left = n16
	n16.Left = n17

	tests := []struct {
		node *binaryTree.Node
		want [][]int
	}{
		{n1, [][]int{{4, 5, 6, 7}, {2, 3}, {1}}},
		{n8, [][]int{{11, 12, 13}, {9, 10}, {8}}},
		{n14, [][]int{{17}, {16}, {15}, {14}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := reverseLevelOrderTraversal(tt.node)
			var gotArr [][]int
			for got.Len() != 0 {
				gotArr = append(gotArr, got.Front().Value.([]int))
				got.Remove(got.Front())
			}

			if !reflect.DeepEqual(gotArr, tt.want) {
				t.Errorf("levelOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}
