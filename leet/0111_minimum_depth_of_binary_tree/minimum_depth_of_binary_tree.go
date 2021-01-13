package _111_minimum_depth_of_binary_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/binaryTree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// tptl? passed.
// todo 2 look for iterative solution
func minDepth(root *binaryTree.Node) int {
	if root != nil {
		left := minDepth(root.Left)
		right := minDepth(root.Right)

		if left == 0 && right == 0 {
			return 1
		}

		if left == 0 {
			return right + 1
		}

		if right == 0 {
			return left + 1
		}

		if left < right {
			return left + 1
		}

		return right + 1
	}

	return 0
}
