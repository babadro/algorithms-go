//https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/discuss/497309/go-solution-O(n)
package _1339_maximum_product_of_splitted_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
)

var total, product int

func maxProduct2(root *binaryTree.Node) int {
	total, product = 0, 0
	total = nodeSum(root)
	total = nodeSum(root)
	return product % (int(1e9) + 7)
}

func nodeSum(root *binaryTree.Node) int {
	if root == nil {
		return 0
	}

	sum := root.Val + nodeSum(root.Left) + nodeSum(root.Right)

	product = max(product, (total-sum)*sum)

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
