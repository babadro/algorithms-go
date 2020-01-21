package _235_lowest_common_ancestor_bst

import (
	bst "github.com/babadro/algorithms-go/leet/0108_sorted_array_to_BST"
	"log"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	arr := []int{6, 2, 8, 0, 4, 7, 9, 3, 5}
	tree := &bst.Tree{}
	for _, n := range arr {
		tree.Insert(&bst.TreeNode{Val: n})
	}
	expected := 6
	lca := lowestCommonAncestor(tree.Root, &bst.TreeNode{Val: 2}, &bst.TreeNode{Val: 8})
	if lca == nil {
		log.Fatalf("lca = nil, expected %d", 6)
	}
	if lca.Val != expected {
		t.Errorf("want %d, got %d", expected, lca.Val)
	}
}
