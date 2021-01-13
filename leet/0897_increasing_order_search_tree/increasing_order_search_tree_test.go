package _897_increasing_order_search_tree

import (
	"fmt"
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/binaryTree"
	"testing"
)

func Test_increasingBST(t *testing.T) {

	input := []int{5, 3, 6, 2, 4, binaryTree.Null, 8, 1, binaryTree.Null, binaryTree.Null, binaryTree.Null, 7, 9}
	root := binaryTree.ArrayToBinaryTree(input)
	got := increasingBST2(root)
	inorderTraversal(got)

}

func inorderTraversal(root *binaryTree.Node) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Println(root.Val)
		inorderTraversal(root.Right)
	}
}
