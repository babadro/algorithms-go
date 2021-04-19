package _429_n_ary_tree_level_order_traversal

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

// passed. dyx.
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	q := list.New()
	q.PushBack(root)

	var res [][]int
	for q.Len() > 0 {
		n := q.Len()

		var level []int
		for i := 0; i < n; i++ {
			node := q.Front().Value.(*Node)
			q.Remove(q.Front())

			level = append(level, node.Val)

			for _, child := range node.Children {
				q.PushBack(child)
			}
		}

		if len(level) > 0 {
			res = append(res, level)
		}
	}

	return res
}

func levelOrderArray(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		level := []int{}
		nextQueue := []*Node{}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			nextQueue = append(nextQueue, node.Children...)
		}
		queue = nextQueue
		result = append(result, level)
	}
	return result
}
