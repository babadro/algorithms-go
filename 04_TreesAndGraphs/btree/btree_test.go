package btree

import (
	"reflect"
	"testing"
)
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

func TestPreOrderFunc(t *testing.T) {
	tests := []struct {
		name        string
		treeBuilder func() *Node
	}{
		{
			name: "1",
			treeBuilder: func() *Node {
				n8 := &Node{Val: 8}
				n3 := &Node{Val: 3}
				n10 := &Node{Val: 10}
				n1 := &Node{Val: 1}
				n6 := &Node{Val: 6}
				n14 := &Node{Val: 14}
				n4 := &Node{Val: 4}
				n7 := &Node{Val: 7}
				n13 := &Node{Val: 13}

				n8.Left, n8.Right = n3, n10
				n3.Left, n3.Right = n1, n6
				n10.Right = n14
				n6.Left, n6.Right = n4, n7
				n14.Left = n13

				return n8
			},
		},
		{
			name: "2",
			treeBuilder: func() *Node {
				n8 := &Node{Val: 8}
				n3 := &Node{Val: 3}
				n10 := &Node{Val: 10}
				n12 := &Node{Val: 12}
				n1 := &Node{Val: 1}
				n0 := &Node{Val: 0}
				n_1 := &Node{Val: -1}
				n6 := &Node{Val: 6}
				n14 := &Node{Val: 14}
				n15 := &Node{Val: 15}
				n4 := &Node{Val: 4}
				n7 := &Node{Val: 7}
				n13 := &Node{Val: 13}
				n9 := &Node{Val: 9}
				n11 := &Node{Val: 11}

				n8.Left, n8.Right = n3, n10
				n3.Left, n3.Right = n1, n6
				n10.Left, n10.Right = n12, n14
				n1.Left, n1.Right = n0, n_1
				n6.Left, n6.Right = n4, n7
				n14.Left, n14.Right = n13, n15
				n12.Left, n12.Right = n9, n11

				return n8
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := make([]int, 0)
			f := func(node *Node) {
				list = append(list, node.Val)
			}

			PreOrderFunc(tt.treeBuilder(), f)

			fmt.Println(list)
		})
	}
}

func TestClone(t *testing.T) {
	tests := []struct {
		name        string
		treeBuilder func() *Node
	}{
		{
			name: "1",
			treeBuilder: func() *Node {
				n8 := &Node{Val: 8}
				n3 := &Node{Val: 3}
				n10 := &Node{Val: 10}
				n1 := &Node{Val: 1}
				n6 := &Node{Val: 6}
				n14 := &Node{Val: 14}
				n4 := &Node{Val: 4}
				n7 := &Node{Val: 7}
				n13 := &Node{Val: 13}

				n8.Left, n8.Right = n3, n10
				n3.Left, n3.Right = n1, n6
				n10.Right = n14
				n6.Left, n6.Right = n4, n7
				n14.Left = n13

				return n8
			},
		},
		{
			name: "2",
			treeBuilder: func() *Node {
				n8 := &Node{Val: 8}
				n3 := &Node{Val: 3}
				n10 := &Node{Val: 10}
				n12 := &Node{Val: 12}
				n1 := &Node{Val: 1}
				n0 := &Node{Val: 0}
				n_1 := &Node{Val: -1}
				n6 := &Node{Val: 6}
				n14 := &Node{Val: 14}
				n15 := &Node{Val: 15}
				n4 := &Node{Val: 4}
				n7 := &Node{Val: 7}
				n13 := &Node{Val: 13}
				n9 := &Node{Val: 9}
				n11 := &Node{Val: 11}

				n8.Left, n8.Right = n3, n10
				n3.Left, n3.Right = n1, n6
				n10.Left, n10.Right = n12, n14
				n1.Left, n1.Right = n0, n_1
				n6.Left, n6.Right = n4, n7
				n14.Left, n14.Right = n13, n15
				n12.Left, n12.Right = n9, n11

				return n8
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := tt.treeBuilder()
			clone := Clone(original)

			if !reflect.DeepEqual(original, clone) {
				originalList := make([]int, 0)
				fOriginal := func(node *Node) {
					originalList = append(originalList, node.Val)
				}
				InOrderFunc(original, fOriginal)

				clonedList := make([]int, 0)
				fCloned := func(node *Node) {
					clonedList = append(clonedList, node.Val)
				}
				InOrderFunc(clone, fCloned)

				t.Errorf("got = %v, want %v", originalList, clonedList)
			}
		})
	}
}
