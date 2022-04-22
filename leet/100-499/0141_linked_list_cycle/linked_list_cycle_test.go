package _141_linked_list_cycle

import (
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
)

func TestHasCycle(t *testing.T) {
	cases := []struct {
		nodes []int
		pos   int
	}{
		{[]int{}, -1},
		{[]int{3, 2, 0, -4}, 1},
		{[]int{1, 2}, 0},
		{[]int{1}, -1},
		{[]int{1}, 0},
		{[]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}, -1},
	}

	for i, c := range cases {
		head := single.ArrToLinkedList(c.nodes)
		var nodeTo, node *single.ListNode
		node = head
		for i := 0; i < len(c.nodes); i++ {
			if nodeTo == nil && i == c.pos {
				nodeTo = node
			}
			if i == len(c.nodes)-1 {
				node.Next = nodeTo
			}
			node = node.Next
		}
		expected := c.pos > -1
		if fact := hasCycleTwoPointers(head); fact != expected {
			t.Errorf("case#%d, want %t, got %t", i+1, expected, fact)
		}
	}
}
