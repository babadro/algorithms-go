package _102_binary_tree_level_order_traversal

import (
	"container/list"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

// passed. tptl. 100% 47%
func levelOrder(root *binaryTree.Node) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	q := list.New()
	q.PushBack(root)

	for q.Len() != 0 {
		level := make([]int, q.Len())
		for i := range level {
			node := q.Front().Value.(*binaryTree.Node)
			q.Remove(q.Front())

			level[i] = node.Val

			if node.Left != nil {
				q.PushBack(node.Left)
			}

			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
