package _094_binary_tree_inorder_traversal

import (
	"fmt"
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tree := btree.ArrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, 7})
	tree2 := btree.ArrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, 7, btree.Null, btree.Null, 8, 9, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 10})
	fmt.Println(inorderTraversalRecursive(tree))
	fmt.Println(inorderTraversalRecursive(tree2))
	//fmt.Println(inorderTraversal(tree))
}
