package _1932_merge_bsts_to_create_single_bst

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"math"
)

// passed. tptl. hard. #tree
// https://leetcode.com/problems/merge-bsts-to-create-single-bst/discuss/1330387/Just-do-what-it-says
func canMerge(trees []*binaryTree.Node) *binaryTree.Node {
	roots := make(map[int]*binaryTree.Node)
	rootCounts := make(map[int]int)

	for _, node := range trees {
		roots[node.Val] = node
		countElements(node, rootCounts)
	}

	var root *binaryTree.Node
	for _, node := range trees {
		if rootCounts[node.Val] == 1 {
			root = node
			delete(roots, node.Val)

			break
		}
	}

	if root == nil {
		return nil
	}

	node := root
	q := []*binaryTree.Node{node}
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if cur.Left == nil && cur.Right == nil && roots[cur.Val] != nil {
			matchingRoot := roots[cur.Val]
			cur.Left, cur.Right = matchingRoot.Left, matchingRoot.Right
			delete(roots, cur.Val)
		}

		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}

	if len(roots) == 0 && isValidBST(root) {
		return root
	} else {
		return nil
	}
}

func countElements(node *binaryTree.Node, m map[int]int) {
	m[node.Val]++
	if node.Left != nil {
		m[node.Left.Val]++
	}
	if node.Right != nil {
		m[node.Right.Val]++
	}
}

func isValidBST(root *binaryTree.Node) bool {
	return recursive(root, math.MinInt64, math.MaxInt64)
}

func recursive(node *binaryTree.Node, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}

	return recursive(node.Left, min, node.Val) && recursive(node.Right, node.Val, max)
}
