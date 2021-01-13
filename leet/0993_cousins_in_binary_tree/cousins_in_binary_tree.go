package _993_cousins_in_binary_tree

import "github.com/babadro/algorithms-go/base/binaryTree"

type data struct {
	x, y             int
	xParent, yParent int
	xDepth, yDepth   int
	isCousins        bool
}

// passed
// todo 2 look for iterative solution
func isCousins(root *binaryTree.Node, x int, y int) bool {
	d := data{x: x, y: y, xDepth: -1, yDepth: -1}

	preOrder(root, nil, 0, &d)

	return d.isCousins
}

func preOrder(root, parent *binaryTree.Node, depth int, data *data) {
	if root == nil {
		return
	}

	if root.Val == data.x {
		if parent == nil {
			return
		}

		data.xParent = parent.Val
		data.xDepth = depth
	} else if root.Val == data.y {
		if parent == nil {
			return
		}

		data.yParent = parent.Val
		data.yDepth = depth
	}

	if data.xDepth != -1 && data.xDepth == data.yDepth {
		if data.xParent == data.yParent {
			data.isCousins = false
		} else {
			data.isCousins = true
		}

		return
	}

	preOrder(root.Left, root, depth+1, data)
	preOrder(root.Right, root, depth+1, data)
}
