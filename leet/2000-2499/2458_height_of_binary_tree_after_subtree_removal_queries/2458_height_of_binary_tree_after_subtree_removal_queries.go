package _2458_height_of_binary_tree_after_subtree_removal_queries

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
)

// #bnsrg
// tle
func treeQueries(root *binaryTree.Node, queries []int) []int {
	res := make([]int, len(queries))

	for i := range res {
		res[i] = height(root, queries[i])
	}

	return res
}

func height(node *binaryTree.Node, removedNode int) int {
	if node == nil || node.Val == removedNode {
		return 0
	}

	if (node.Left == nil || node.Left.Val == removedNode) &&
		(node.Right == nil || node.Right.Val == removedNode) {
		return 0
	}

	return 1 + max(height(node.Left, removedNode), height(node.Right, removedNode))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
