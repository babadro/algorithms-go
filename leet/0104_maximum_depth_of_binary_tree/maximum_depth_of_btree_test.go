package _104_maximum_depth_of_binary_tree

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := btree.ArrayToBinaryTree([]int{1, 5, 2, 3, 4, btree.Null, btree.Null, btree.Null, btree.Null, 6, 7})

	t.Log(maxDepth(root))
}
