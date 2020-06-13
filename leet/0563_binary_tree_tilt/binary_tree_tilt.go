package _563_binary_tree_tilt

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
)

func findTilt(root *btree.Node) int {
	tiltSum := 0
	postOrder(root, &tiltSum)
	return tiltSum
}

func postOrder(node *btree.Node, tiltSum *int) (sumValues int) {
	if node != nil {
		leftSumValues := postOrder(node.Left, tiltSum)
		rightSumValues := postOrder(node.Right, tiltSum)
		*tiltSum += mod(leftSumValues, rightSumValues)
		return leftSumValues + rightSumValues + node.Val
	}
	return 0
}

func mod(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}
