package _111_minimum_depth_of_binary_tree

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func Test_minDepth(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"1", []int{3, 9, 20, btree.Null, btree.Null, 15, 7}, 2},
		//{"2", []int{2, 3, btree.Null, 4, btree.Null, btree.Null, btree.Null, 5, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.ArrayToBinaryTree(tt.arr)
			if got := minDepth(root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepth2(t *testing.T) {
	n2 := &btree.Node{Val: 2}
	n3 := &btree.Node{Val: 3}
	n4 := &btree.Node{Val: 4}
	n5 := &btree.Node{Val: 5}
	n6 := &btree.Node{Val: 6}

	n2.Left = n3
	n3.Left = n4
	n4.Left = n5
	n5.Left = n6

	t.Log(minDepth(n2))
}
