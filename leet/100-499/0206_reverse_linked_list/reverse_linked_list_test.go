package _206_reverse_linked_list

import (
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
	"github.com/babadro/algorithms-go/slices"
)

func TestReverseList(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{0, 1}, []int{1, 0}},
		{[]int{0, 1, 2, 3}, []int{3, 2, 1, 0}},
	}

	for i, c := range cases {
		head := singly.ArrToLinkedList(c.input)
		reverseList := reverseList(head)
		fact := singly.LinkedListToArr(reverseList)
		if !slices.IntSlicesAreEqual(fact, c.expected) {
			t.Errorf("case#%d, want %v, got %v", i+1, c.expected, fact)
		}
	}
}
