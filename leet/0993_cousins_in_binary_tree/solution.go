package _993_cousins_in_binary_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

// best solution. passed. tptl
func isCousins2(root *btree.Node, x, y int) bool {
	xDepth, yDepth, xParent, yParent := 0, 0, 0, 0
	walker(root, x, y, 0, 0, &xDepth, &yDepth, &xParent, &yParent)

	return xDepth == yDepth && xParent != yParent
}

func walker(node *btree.Node, x, y, p, depth int, xDepth, yDepth, xParent, yParent *int) {
	if node == nil {
		return
	}

	if node.Val == x {
		*xDepth = depth
		*xParent = p
	}

	if node.Val == y {
		*yDepth = depth
		*yParent = p
	}

	walker(node.Left, x, y, node.Val, depth+1, xDepth, yDepth, xParent, yParent)
	walker(node.Right, x, y, node.Val, depth+1, xDepth, yDepth, xParent, yParent)
}
