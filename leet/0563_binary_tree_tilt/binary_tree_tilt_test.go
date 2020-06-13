package _563_binary_tree_tilt

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func TestFindTilt(t *testing.T) {
	root := btree.ArrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, 7})
	t.Log(findTilt(root))
}

func TestPostOrder(t *testing.T) {
	root := btree.ArrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, 7})
	tiltSum := 0
	valuesSum := postOrder(root, &tiltSum)
	t.Log("tiltSum: ", tiltSum, " valuesSum: ", valuesSum)
}
