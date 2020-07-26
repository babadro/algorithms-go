package _098_validate_binary_search_tree

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		arr      []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{1, 2, 3}, true},
	}

	for i, c := range cases {
		root := btree.ArrayToBinaryTree(c.arr)
		if fact := isValidBST(root); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
