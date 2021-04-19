package _235_lowest_common_ancestor_bst

import (
	"github.com/babadro/algorithms-go/leet/100-499/0108_sorted_array_to_BST"
	"log"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	arr := []int{6, 2, 8, 0, 4, 7, 9, 3, 5}
	tree := &_0108_sorted_array_to_BST.Tree{}
	for _, n := range arr {
		tree.Insert(&_0108_sorted_array_to_BST.TreeNode{Val: n})
	}
	expected := 6
	lca := lowestCommonAncestor(tree.Root, &_0108_sorted_array_to_BST.TreeNode{Val: 2}, &_0108_sorted_array_to_BST.TreeNode{Val: 8})
	if lca == nil {
		log.Fatalf("lca = nil, expected %d", 6)
	}
	if lca.Val != expected {
		t.Errorf("want %d, got %d", expected, lca.Val)
	}
}
