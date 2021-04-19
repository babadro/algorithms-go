package _112_path_sum

import "github.com/babadro/algorithms-go/base/binaryTree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// passed.
// todo 3 look for iterative solution
func hasPathSum(root *binaryTree.Node, sum int) bool {
	found := false
	helper(root, 0, sum, &found)

	return found
}

func helper(node *binaryTree.Node, curSum, target int, found *bool) {
	if *found || node == nil {
		return
	}

	curSum += node.Val

	if node.Left == nil && node.Right == nil && curSum == target {
		*found = true

		return
	}

	helper(node.Left, curSum, target, found)
	helper(node.Right, curSum, target, found)
}
