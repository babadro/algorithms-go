package _104_maximum_depth_of_binary_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/binaryTree"

func maxDepth(root *binaryTree.Node) int {
	maxD := 0
	preOrder(root, 0, &maxD)
	return maxD
}

func preOrder(node *binaryTree.Node, level int, maxD *int) {
	if node != nil {
		level++
		if level > *maxD {
			*maxD = level
		}
		preOrder(node.Left, level, maxD)
		preOrder(node.Right, level, maxD)
	}
}

/* code for submission
func maxDepth(root *TreeNode) int {
	maxD := 0
	preOrder(root, 0, &maxD)
	return maxD
}

func preOrder(node *TreeNode, level int, maxD *int) {
	if node != nil {
		level++
		if level > *maxD {
			*maxD = level
		}
		preOrder(node.Left, level, maxD)
		preOrder(node.Right, level, maxD)
	}
}

*/
