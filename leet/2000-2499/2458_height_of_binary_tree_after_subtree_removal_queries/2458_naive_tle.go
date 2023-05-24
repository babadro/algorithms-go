package _2458_height_of_binary_tree_after_subtree_removal_queries

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"github.com/babadro/algorithms-go/utils"
)

func treeQueriesTLE(root *binaryTree.Node, queries []int) []int {
	res := make([]int, len(queries))

	for i := range res {
		res[i] = heightTLE(root, queries[i])
	}

	return res
}

func heightTLE(node *binaryTree.Node, removedNode int) int {
	if node == nil || node.Val == removedNode {
		return 0
	}

	if (node.Left == nil || node.Left.Val == removedNode) &&
		(node.Right == nil || node.Right.Val == removedNode) {
		return 0
	}

	return 1 + utils.Max(heightTLE(node.Left, removedNode), heightTLE(node.Right, removedNode))
}
