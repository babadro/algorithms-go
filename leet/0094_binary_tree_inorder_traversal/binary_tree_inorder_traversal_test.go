package _094_binary_tree_inorder_traversal

import (
	"fmt"
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tree := btree.ArrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(inorderTraversalRecursive(tree))
}
