package _098_validate_binary_search_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"math"
)

// 97% 80%
func isValidBST(root *binaryTree.Node) bool {
	return recursive(root, math.MinInt64, math.MaxInt64)
}

func recursive(node *binaryTree.Node, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}

	return recursive(node.Left, min, node.Val) && recursive(node.Right, node.Val, max)
}
