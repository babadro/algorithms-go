package _1_binary_tree_level_order_traversal

import (
	"container/list"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

// leetcode 102
func levelOrderTraversal(node *binaryTree.Node) [][]int {
	if node == nil {
		return nil
	}
	var res [][]int
	q := list.New()
	q.PushBack(node)

	for q.Len() != 0 {
		level := make([]int, q.Len())

		for i := 0; i < len(level); i++ {
			n := q.Front().Value.(*binaryTree.Node)
			q.Remove(q.Front())

			level[i] = n.Val

			if n.Left != nil {
				q.PushBack(n.Left)
			}

			if n.Right != nil {
				q.PushBack(n.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
