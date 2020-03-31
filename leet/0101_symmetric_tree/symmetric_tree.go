package _101_symmetric_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

type item struct {
	val   int
	level int
}

func isSymmetric(root *btree.Node) bool {
	arr := make([]item, 0)
	rootAchieved := false
	var idx int
	res := true
	f := func(level int, n *btree.Node) {
		if n == root {
			rootAchieved = true
			idx = len(arr) - 1
			return
		}
		if rootAchieved {
			if idx < 0 || arr[idx].val != n.Val || arr[idx].level != level {
				res = false
			}
			idx--
		} else {
			arr = append(arr, item{n.Val, level})
		}
	}
	inOrder(0, root, f)
	return res
}

func inOrder(level int, root *btree.Node, f func(level int, n *btree.Node)) {
	if root != nil {
		inOrder(level+1, root.Left, f)
		f(level, root)
		inOrder(level+1, root.Right, f)
	}
}
