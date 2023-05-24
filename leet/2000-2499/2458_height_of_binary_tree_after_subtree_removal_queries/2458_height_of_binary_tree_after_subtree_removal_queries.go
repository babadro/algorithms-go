package _2458_height_of_binary_tree_after_subtree_removal_queries

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
)

type pathNode struct {
	val              int
	height           int
	smallChildHeight int
}

// #bnsrg 62%, 50%
// todo 2 find shorter and if possible faster solution
func treeQueries(root *binaryTree.Node, queries []int) []int {
	nodes := heights(root)

	treeHeightAfterRemoval := make(map[int]int, len(nodes))
	maxH, level := 0, 0
	for i := len(nodes) - 2; i >= 0; i-- {
		parent := nodes[i+1]
		maxH = max(maxH, level+parent.smallChildHeight)
		treeHeightAfterRemoval[nodes[i].val] = maxH

		level++
	}

	res := make([]int, len(queries))
	rootHeight := nodes[len(nodes)-1].height - 1

	for i := range res {
		if h, ok := treeHeightAfterRemoval[queries[i]]; ok {
			res[i] = h
		} else {
			res[i] = rootHeight
		}
	}

	return res
}

func heights(n *binaryTree.Node) []pathNode {
	if n == nil {
		return []pathNode{{val: -1, height: 0, smallChildHeight: 0}}
	}

	high, low := heights(n.Left), heights(n.Right)

	highH, lowH := high[len(high)-1].height, low[len(low)-1].height

	if highH < lowH {
		high, low = low, high
		highH, lowH = lowH, highH
	}

	resNode := pathNode{
		val:              n.Val,
		height:           highH + 1,
		smallChildHeight: lowH,
	}

	if highH == lowH {
		return []pathNode{resNode}
	}

	return append(high, resNode)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
