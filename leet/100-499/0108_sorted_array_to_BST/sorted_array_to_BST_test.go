package _0108_sorted_array_to_BST

import (
	"fmt"
	"testing"
)

// Just for fun
func TestPreOrder(t *testing.T) {
	arr := []int{6, 2, 8, 0, 4, 7, 9, 3, 5}
	bst := &Tree{}
	for _, n := range arr {
		bst.Insert(&TreeNode{Val: n})
	}
	list := make([]int, 0)
	PreOrder(bst.Root, &list)
	for _, n := range list {
		fmt.Println(n)
	}
}

func TestLevelOrder(t *testing.T) {
	arr := []int{6, 2, 8, 0, 4, 7, 9, 3, 5}
	bst := &Tree{}
	for _, n := range arr {
		bst.Insert(&TreeNode{Val: n})
	}
	list := LevelOrder(bst.Root)
	for _, n := range list {
		fmt.Println(n)
	}
}
