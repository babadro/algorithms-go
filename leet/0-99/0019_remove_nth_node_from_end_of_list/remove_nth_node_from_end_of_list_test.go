package _019_remove_nth_node_from_end_of_list

import (
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
	"github.com/babadro/algorithms-go/slices"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		input    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{1, 3, 4, 5}},
	}

	for i, c := range cases {
		head := singly.ArrToLinkedList(c.input)
		head = removeNthFromEnd(head, c.n)
		fact := singly.LinkedListToArr(head)
		if !slices.IntSlicesAreEqual(fact, c.expected) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, fact)
		}
	}
}
