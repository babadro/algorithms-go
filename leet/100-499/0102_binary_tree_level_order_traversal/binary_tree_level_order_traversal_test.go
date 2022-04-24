package _102_binary_tree_level_order_traversal

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func TestLevelOrder(t *testing.T) {
	root := binaryTree.ArrayToBinaryTree([]int{3, 9, 20, binaryTree.Null, binaryTree.Null, 15, 7})
	t.Log(levelOrder(root))
	root2 := &binaryTree.Node{Val: 1}
	n2 := &binaryTree.Node{Val: 2}
	n3 := &binaryTree.Node{Val: 3}
	n4 := &binaryTree.Node{Val: 4}
	root2.Right = n2
	n2.Right = n3
	n3.Right = n4
	t.Log(levelOrder(root2))
}
