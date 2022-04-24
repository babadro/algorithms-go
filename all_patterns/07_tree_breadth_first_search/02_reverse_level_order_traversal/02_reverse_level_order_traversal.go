package _2_reverse_level_order_traversal

import (
	"container/list"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

// not interesting. It is just returning linkedList of []int instead of [][]int as a result
// I could reverse result [][]int array as well instead of returning *list.List
func reverseLevelOrderTraversal(node *binaryTree.Node) *list.List {
	if node == nil {
		return nil
	}

	res, queue := list.New(), list.New()
	queue.PushBack(node)

	for queue.Len() != 0 {
		level := make([]int, queue.Len())
		for i := 0; i < len(level); i++ {
			n := queue.Front().Value.(*binaryTree.Node)
			queue.Remove(queue.Front())

			level[i] = n.Val

			if n.Left != nil {
				queue.PushBack(n.Left)
			}

			if n.Right != nil {
				queue.PushBack(n.Right)
			}
		}

		res.PushFront(level)
	}

	return res
}
