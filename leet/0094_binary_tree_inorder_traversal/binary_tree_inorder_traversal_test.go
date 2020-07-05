package _094_binary_tree_inorder_traversal

import (
	"fmt"
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tree := btree.ArrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, 7})
	//tree2 := btree.ArrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, 7, btree.Null, btree.Null, 8, 9, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, btree.Null, 10})
	res1 := make([]int, 0)
	inorderTraversalRecursive(tree, &res1)
	fmt.Println(res1)
	var res2 []int
	tree2 := btree.ArrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	inorderTraversalRecursive(tree2, &res2)
	fmt.Println(res2)
	n8, n9, n10, n11, n12, n13, n14, n15 := &btree.Node{Val: 8}, &btree.Node{Val: 9}, &btree.Node{Val: 10}, &btree.Node{Val: 11}, &btree.Node{Val: 12}, &btree.Node{Val: 13}, &btree.Node{Val: 14}, &btree.Node{Val: 15}
	n4, n5, n6, n7 := &btree.Node{Val: 4}, &btree.Node{Val: 5}, &btree.Node{Val: 6}, &btree.Node{Val: 7}
	n2, n3 := &btree.Node{Val: 2}, &btree.Node{Val: 3}
	n1 := &btree.Node{Val: 1}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right, n3.Left, n3.Right = n4, n5, n6, n7
	n4.Left, n4.Right, n5.Left, n5.Right, n6.Left, n6.Right, n7.Left, n7.Right = n8, n9, n10, n11, n12, n13, n14, n15
	res3 := make([]int, 0)
	inorderTraversalRecursive(n1, &res3)
	fmt.Println(res3)
	//fmt.Println(inorderTraversal(tree))
}

func TestInorderTraversal2(t *testing.T) {
	n10 := &btree.Node{Val: 10}
	n8, n9 := &btree.Node{Val: 8}, &btree.Node{Val: 9}
	n4, n5, n6, n7 := &btree.Node{Val: 4}, &btree.Node{Val: 5}, &btree.Node{Val: 6}, &btree.Node{Val: 7}
	n2, n3 := &btree.Node{Val: 2}, &btree.Node{Val: 3}
	n1 := &btree.Node{Val: 1}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right, n3.Left, n3.Right = n4, n5, n6, n7
	n5.Left, n5.Right = n8, n9
	n9.Right = n10

	var res []int
	inorderTraversalRecursive(n1, &res)
	fmt.Println(res)
}
