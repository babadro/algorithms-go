package _107_binary_tree_level_rder_traversal_ii

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_levelOrderBottom(t *testing.T) {
	n3, n9, n20, n15, n7 := n(3), n(9), n(20), n(15), n(7)

	n3.Left, n3.Right = n9, n20
	n20.Left, n20.Right = n15, n7

	t.Log(levelOrderBottom(n3))
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}
