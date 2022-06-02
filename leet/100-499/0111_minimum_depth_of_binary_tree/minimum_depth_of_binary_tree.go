package _111_minimum_depth_of_binary_tree

import (
	"math"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

// iterative.
// [1,2,3,4,5] test case failed
func minDepth3(root *binaryTree.Node) int {
	if root == nil {
		return 0
	}

	q := []*binaryTree.Node{root}
	res := 0
	for len(q) > 0 {
		levelSize := len(q)
		res++
		var node *binaryTree.Node
		for i := 0; i < levelSize; i++ {
			node, q = q[0], q[1:]
			if node.Left == nil && node.Right == nil {
				return res
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return res
}

// passed. iterative. easy to understand
func minDepth2(root *binaryTree.Node) int {
	if root == nil {
		return 0
	}

	res := math.MaxInt64
	rec(root, 1, &res)

	return res
}

func rec(node *binaryTree.Node, cur int, res *int) {
	if cur >= *res {
		return
	}

	if node.Left == nil && node.Right == nil {
		*res = min(*res, cur)
	}

	if node.Left != nil {
		rec(node.Left, cur+1, res)
	}

	if node.Right != nil {
		rec(node.Right, cur+1, res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//  passed.
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
