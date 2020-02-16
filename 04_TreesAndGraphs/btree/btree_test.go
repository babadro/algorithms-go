package btree

import "testing"
import "fmt"

func TestArrayToBinaryTree(t *testing.T) {
	cases := []struct {
		arr []int
	}{
		{[]int{1, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 5, 2, 7, Null, 3, 4, Null, Null, Null, Null, 5, 6}},
		{[]int{1, Null, 2, 3, 4, Null, Null, 5, 6}},
		{[]int{1, 5, 2, 3, 4, Null, Null, Null, Null, 6, 7}},
		{[]int{2, 3, 9, 10, 7, 8, 6, 5, 4, 11, 1}},
	}

	for _, c := range cases {
		root := ArrayToBinaryTree(c.arr)
		list := new([]int)
		InOrder(root, list)
		t.Log(list)
	}
}

func TestInorderFunc(t *testing.T) {
	cases := []struct {
		arr []int
	}{
		{[]int{1, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 5, 2, 7, Null, 3, 4, Null, Null, Null, Null, 5, 6}},
		{[]int{1, Null, 2, 3, 4, Null, Null, 5, 6}},
		{[]int{1, 5, 2, 3, 4, Null, Null, Null, Null, 6, 7}},
		{[]int{2, 3, 9, 10, 7, 8, 6, 5, 4, 11, 1}},
	}

	for _, c := range cases {
		root := ArrayToBinaryTree(c.arr)
		list := make([]int, 0)
		f := func(node *Node) {
			list = append(list, node.Val)
		}
		InOrderFunc(root, f)
		fmt.Println(list)
	}

}
