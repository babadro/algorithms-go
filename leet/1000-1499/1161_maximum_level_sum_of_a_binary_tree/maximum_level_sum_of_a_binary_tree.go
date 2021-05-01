package _1161_maximum_level_sum_of_a_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
)

// passed. tptl. medium. best solution (not mine)
func maxLevelSum(root *binaryTree.Node) int {
	q, res, max, level := []*binaryTree.Node{root}, 1, root.Val, 1

	for len(q) > 0 {
		sum := 0
		var node *binaryTree.Node
		for l := len(q); l > 0; l-- {
			node, q = q[0], q[1:]
			sum += node.Val

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		if sum > max {
			max, res = sum, level
		}

		level++
	}

	return res
}
