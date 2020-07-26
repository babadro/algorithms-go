package _098_validate_binary_search_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

// TODO 1
func isValidBST(root *btree.Node) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Val <= root.Left.Val {
		return false
	}
	if root.Right != nil && root.Val >= root.Right.Val {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}
