package btree

import "testing"

func TestArrayToBinaryTree(t *testing.T) {
	cases := []struct {
		arr []int
	}{
		{[]int{1, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 5, 2, 7, null, 3, 4, null, null, null, null, null, null, 5, 6}},
		{[]int{1, null, 2, 3, 4, null, null, 5, 6}},
		{[]int{1, 5, 2, 3, 4, null, null, null, null, 6, 7, null, null, 8, 9}},
	}

	for _, c := range cases {
		root := ArrayToBinaryTree(c.arr)
		list := new([]int)
		InOrder(root, list)
		t.Log(list)
	}
}
