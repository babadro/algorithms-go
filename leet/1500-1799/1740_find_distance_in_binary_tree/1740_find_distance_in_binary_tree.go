package _1740_find_distance_in_binary_tree

import "github.com/babadro/algorithms-go/base/binaryTree"

func findDistance(root *binaryTree.Node, p int, q int) int {
	var res int

	traverse(root, p, q, &res)

	return res
}

// #bnsrg medium passed
func traverse(node *binaryTree.Node, p, q int, res *int) int {
	if node == nil {
		return 0
	}

	left, right := traverse(node.Left, p, q, res), traverse(node.Right, p, q, res)

	if node.Val == p || node.Val == q {
		if left != 0 {
			*res = left
		} else if right != 0 {
			*res = right
		}

		return 1
	}

	if left != 0 && right != 0 {
		*res = left + right

		return 0
	}

	if left != 0 {
		return left + 1
	}

	if right != 0 {
		return right + 1
	}

	return 0
}
