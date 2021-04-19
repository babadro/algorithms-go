package _814_binary_tree_pruning

import "github.com/babadro/algorithms-go/base/binaryTree"

// dyx. passed.
func pruneTree(root *binaryTree.Node) *binaryTree.Node {
	if root != nil {
		root.Left = pruneTree(root.Left)
		root.Right = pruneTree(root.Right)

		if root.Val == 0 && root.Left == nil && root.Right == nil {
			return nil
		}
	}

	return root
}
