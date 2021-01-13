package _103_binary_tree_zigzag_level_order_traversal

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/binaryTree"
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {

	tests := []struct {
		name string
		arr  []int
		want [][]int
	}{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "1",
			arr:  []int{1},
			want: [][]int{{1}},
		},
		{
			name: "2",
			arr:  []int{1, 2, 3},
			want: [][]int{{1}, {3, 2}},
		},
		{
			name: "3",
			arr:  []int{1, 2, 3, 4, 5, 6, 7},
			want: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := binaryTree.ArrayToBinaryTree(tt.arr)
			if got := zigzagLevelOrder(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
