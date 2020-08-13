package _102_binary_tree_level_order_traversal

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := btree.ArrayToBinaryTree([]int{3, 9, 20, btree.Null, btree.Null, 15, 7})
	t.Log(levelOrder(root))
	root2 := &btree.Node{Val: 1}
	n2 := &btree.Node{Val: 2}
	n3 := &btree.Node{Val: 3}
	n4 := &btree.Node{Val: 4}
	root2.Right = n2
	n2.Right = n3
	n3.Right = n4
	t.Log(levelOrder(root2))
}
