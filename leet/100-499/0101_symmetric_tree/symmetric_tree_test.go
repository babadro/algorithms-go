package _101_symmetric_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	cases := []struct {
		array    []int
		expected bool
	}{
		//{[]int{}, true},
		//{[]int{1}, true},
		//{[]int{1, 2, 2, binaryTree.Null, 3, binaryTree.Null, 3}, false},
		{[]int{1, 2, 2, binaryTree.Null, 3, 3}, true},
		//{[]int{1, 2, 2}, true},
		//{[]int{1, 2, 2, 2, binaryTree.Null, 2}, false},
		//{[]int{1, 0}, false},
		//{[]int{1, binaryTree.Null, 2}, false},
		//{[]int{2, 3, 3, 4, 5, 5}, false},
	}

	for i, c := range cases {
		tree := binaryTree.ArrayToBinaryTree(c.array)
		if fact := isSymmetric(tree); fact != c.expected {
			t.Errorf("case#%d, want isSymmetric= %t, got isSymmetric= %t", i+1, c.expected, fact)
		}
	}
}
