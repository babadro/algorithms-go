package __path_with_given_sequence

import "github.com/babadro/algorithms-go/base/binaryTree"

// probably leetcode 1430 (premium)
func findPath(root *binaryTree.Node, seq []int) bool {
	return preOrder(root, 0, seq)
}

func preOrder(node *binaryTree.Node, idx int, seq []int) bool {
	if node == nil {
		return idx == len(seq)
	}

	if seq[idx] != node.Val {
		return false
	}

	idx++

	return preOrder(node.Left, idx, seq) || preOrder(node.Right, idx, seq)
}
