package _897_increasing_order_search_tree

import (
	"fmt"
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func Test_increasingBST(t *testing.T) {

	input := []int{5, 3, 6, 2, 4, btree.Null, 8, 1, btree.Null, btree.Null, btree.Null, 7, 9}
	root := btree.ArrayToBinaryTree(input)
	got := increasingBST2(root)
	inorderTraversal(got)

}

func inorderTraversal(root *btree.Node) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Println(root.Val)
		inorderTraversal(root.Right)
	}
}
