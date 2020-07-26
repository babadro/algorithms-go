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
		{[]int{10, 9, 11, btree.Null, btree.Null, 9, 12}, false},
		{[]int{10, 5, 15, 3, 7, 12, 17, 1, 4, 6, 8, 11, 13, 16, 18}, true},
		{[]int{1, 1}, false},
	}

	for i, c := range cases {
		root := btree.ArrayToBinaryTree(c.arr)
		if fact := isValidBST(root); fact != c.expected {
			t.Errorf("case#%d, want %t, got %t", i+1, c.expected, fact)
		}
	}
}

func TestLeftList(t *testing.T) {
	n4 := &btree.Node{Val: 4}
	n3 := &btree.Node{Val: 3}
	n2 := &btree.Node{Val: 2}
	n1 := &btree.Node{Val: 1}

	n4.Left = n3
	n3.Left = n2
	n2.Left = n1

	t.Log(isValidBST(n4))
}

func TestRightList(t *testing.T) {
	n1 := &btree.Node{Val: 1}
	n2 := &btree.Node{Val: 2}
	n3 := &btree.Node{Val: 3}
	n4 := &btree.Node{Val: 4}

	n1.Right = n2
	n2.Right = n3
	n3.Right = n4

	t.Log(isValidBST(n1))
}
