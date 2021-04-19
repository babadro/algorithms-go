package _1339_maximum_product_of_splitted_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
)

var maxProduction int
var sumAllNodes int

func maxProduct(root *binaryTree.Node) int {
	maxProduction, sumAllNodes = 0, 0
	// Get sum all nodes in tree
	sumAllNodes = treeTraversal(root)
	// Calc max possible production (maxProduction) of splitted tree
	_ = treeTraversal(root)

	return maxProduction % (int(1e9) + 7)
}

func treeTraversal(node *binaryTree.Node) int {
	if node == nil {
		return 0
	}
	currSum := node.Val + treeTraversal(node.Left) + treeTraversal(node.Right)
	maxProduction = calcMax(maxProduction, (sumAllNodes-currSum)*currSum)
	return currSum
}

func calcMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
