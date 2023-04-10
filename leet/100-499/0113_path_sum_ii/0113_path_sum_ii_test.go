package _113_path_sum_ii

import (
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_pathSum(t *testing.T) {
	// root1
	n1_5 := n(5)
	n2_4, n2_8 := n(4), n(8)
	n3_11, n3_13, n3_4 := n(11), n(13), n(4)
	n4_7, n4_2, n4_5, n4_1 := n(7), n(2), n(5), n(1)

	n1_5.Left, n1_5.Right = n2_4, n2_8
	n2_4.Left = n3_11
	n2_8.Left, n2_8.Right = n3_13, n3_4
	n3_11.Left, n3_11.Right = n4_7, n4_2
	n3_4.Left, n3_4.Right = n4_5, n4_1

	// root2
	n1, n2, n3 := n(1), n(2), n(3)
	n1.Left, n1.Right = n2, n3

	tests := []struct {
		root      *binaryTree.Node
		targetSum int
		want      [][]int
	}{
		{n1_5, 22, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		{n1, 5, nil},
		{nil, 5, nil},
		{n(1), 1, [][]int{{1}}},
		{n(1), 2, nil},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := pathSum2(tt.root, tt.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(v int) *binaryTree.Node {
	return &binaryTree.Node{Val: v}
}
