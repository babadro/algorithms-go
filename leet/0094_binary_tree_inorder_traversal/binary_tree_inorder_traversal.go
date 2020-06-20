package _094_binary_tree_inorder_traversal

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

var res []int

// TODO 1 - iterative
func inorderTraversal(root *btree.Node) []int {

}

func inorderTraversalRecursive(root *btree.Node) []int {
	if root != nil {
		inorderTraversalRecursive(root.Left)
		res = append(res, root.Val)
		inorderTraversalRecursive(root.Right)
	}
	return res
}
