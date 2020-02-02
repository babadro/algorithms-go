package btree

import "math"

const (
	null = math.MaxInt64
)

func ArrayToBinaryTree(arr []int) (root *TreeNode) {

	var i int
	currLevel := uint(0)
	level := make([]*TreeNode, 0)
	for i < len(arr) {
		for k, n := range level {
			i++
			n.Left
		}
		level = level[:0]
		for j := 0; j < 1<<currLevel; j++ {

			level = append(level, node)
		}

		node.Val = arr[i]
		for j, v := range
	}
}