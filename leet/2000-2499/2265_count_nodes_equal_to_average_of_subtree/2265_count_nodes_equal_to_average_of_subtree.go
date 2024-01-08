package _2265_count_nodes_equal_to_average_of_subtree

import "github.com/babadro/algorithms-go/base/binaryTree"

// bnsrg passed medium
func averageOfSubtree(root *binaryTree.Node) int {
	res := 0

	traverse(root, &res)

	return res
}

func traverse(n *binaryTree.Node, counter *int) (int, int) {
	if n == nil {
		return 0, 0
	}

	leftCount, leftSum := traverse(n.Left, counter)
	rightCount, rightSum := traverse(n.Right, counter)

	count := leftCount + rightCount + 1
	sum := leftSum + rightSum + n.Val

	if sum/count == n.Val {
		*counter++
	}

	return count, sum
}
