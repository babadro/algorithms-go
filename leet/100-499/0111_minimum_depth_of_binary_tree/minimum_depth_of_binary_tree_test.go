package _111_minimum_depth_of_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"testing"
)

func Test_minDepth(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"1", []int{3, 9, 20, binaryTree.Null, binaryTree.Null, 15, 7}, 2},
		//{"2", []int{2, 3, binaryTree.Null, 4, binaryTree.Null, binaryTree.Null, binaryTree.Null, 5, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := binaryTree.ArrayToBinaryTree(tt.arr)
			if got := minDepth(root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepth2(t *testing.T) {
	n2 := &binaryTree.Node{Val: 2}
	n3 := &binaryTree.Node{Val: 3}
	n4 := &binaryTree.Node{Val: 4}
	n5 := &binaryTree.Node{Val: 5}
	n6 := &binaryTree.Node{Val: 6}

	n2.Left = n3
	n3.Left = n4
	n4.Left = n5
	n5.Left = n6

	t.Log(minDepth(n2))
}
