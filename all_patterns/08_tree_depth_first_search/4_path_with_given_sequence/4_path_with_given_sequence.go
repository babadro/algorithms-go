package __path_with_given_sequence

import "github.com/babadro/algorithms-go/base/binaryTree"

// probably leetcode 1430 (premium)
// Check If a String Is a Valid Sequence from Root
// to Leaves Path in a Binary Tree
// https://leetcode.ca/all/1430.html
func findPath(node *binaryTree.Node, seq []int) bool {
	if node == nil || len(seq) == 0 || node.Val != seq[0] {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return len(seq) == 1
	}

	return findPath(node.Left, seq[1:]) || findPath(node.Right, seq[1:])
}
