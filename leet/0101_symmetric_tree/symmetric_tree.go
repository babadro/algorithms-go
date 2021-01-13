package _101_symmetric_tree

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/binaryTree"
)

type item struct {
	val   int
	level int
}

func isSymmetric(root *binaryTree.Node) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left, right *binaryTree.Node) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) || (left.Val != right.Val) {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

func isSymmetric2(root *binaryTree.Node) bool {
	leftBranch := make([]item, 0)
	rootAchieved := false
	var idx int
	flag := true
	lenRightBranch := 0
	f := func(level int, n *binaryTree.Node) {
		if n == root {
			rootAchieved = true
			idx = len(leftBranch) - 1
			return
		}
		if rootAchieved {
			if idx < 0 || leftBranch[idx].val != n.Val || leftBranch[idx].level != level {
				flag = false
			}
			lenRightBranch++
			idx--
		} else {
			leftBranch = append(leftBranch, item{n.Val, level})
		}
	}
	inOrder(0, root, f)
	if len(leftBranch) != lenRightBranch {
		return false
	}
	if len(leftBranch) == 0 {
		return true
	}
	return flag
}

func inOrder(level int, root *binaryTree.Node, f func(level int, n *binaryTree.Node)) {
	if root != nil {
		inOrder(level+1, root.Left, f)
		f(level, root)
		inOrder(level+1, root.Right, f)
	}
}
