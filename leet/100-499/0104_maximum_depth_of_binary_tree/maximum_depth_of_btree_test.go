package _104_maximum_depth_of_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := binaryTree.ArrayToBinaryTree([]int{1, 5, 2, 3, 4, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, 6, 7})

	t.Log(maxDepth(root))
}
