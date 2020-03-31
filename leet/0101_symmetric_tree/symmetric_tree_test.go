package _101_symmetric_tree

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	cases := []struct {
		array    []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{1, 2, 2, btree.Null, 3, btree.Null, 3}, false},
		{[]int{1, 2, 2, btree.Null, 3, 3}, true},
		{[]int{1, 2, 2}, true},
		{[]int{1, 2, 2, 2, btree.Null, 2}, false},
		{[]int{1, 0}, false}, // TODO
	}

	for i, c := range cases {
		tree := btree.ArrayToBinaryTree(c.array)
		if fact := isSymmetric(tree); fact != c.expected {
			t.Errorf("case#%d, want isSymmetric= %t, got isSymmetric= %t", i+1, c.expected, fact)
		}
	}
}
