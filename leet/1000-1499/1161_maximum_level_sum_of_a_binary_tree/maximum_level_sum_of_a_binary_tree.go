package _1161_maximum_level_sum_of_a_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"math"
)

type nodeInfo struct {
	level int
	*binaryTree.Node
}

// passed. medium
func maxLevelSum(root *binaryTree.Node) int {
	if root == nil {
		return 1
	}

	queue := []nodeInfo{{1, root}}

	level, max := 0, math.MinInt32
	currLevel, sum := 0, math.MinInt32
	var node nodeInfo
	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]

		if node.level != currLevel {
			if sum > max {
				max, level = sum, currLevel
			}

			currLevel, sum = node.level, node.Val
		} else {
			sum += node.Val
		}

		if node.Left != nil {
			queue = append(queue, nodeInfo{
				level: currLevel + 1,
				Node:  node.Left,
			})
		}

		if node.Right != nil {
			queue = append(queue, nodeInfo{
				level: currLevel + 1,
				Node:  node.Right,
			})
		}
	}

	if sum > max {
		max, level = sum, currLevel
	}

	return level
}
