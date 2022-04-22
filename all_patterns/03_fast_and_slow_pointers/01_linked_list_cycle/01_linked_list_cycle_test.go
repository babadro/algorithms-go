package _1_linked_list_cycle

import (
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
)

func Test_hasCycle(t *testing.T) {
	var nilHead *single.ListNode

	singleNode := n(1)

	n1, n2, n3, n4, n5, n6, n7 := n(1), n(1), n(1), n(1), n(1), n(1), n(1)
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7

	cycle1, cycle2, cycle3, cycle4, cycle5, cycle6 := n(1), n(1), n(1), n(1), n(1), n(1)
	cycle1.Next = cycle2
	cycle2.Next = cycle3
	cycle3.Next = cycle4
	cycle4.Next = cycle5
	cycle5.Next = cycle6
	cycle6.Next = cycle3

	tests := []struct {
		node *single.ListNode
		want bool
	}{
		{nilHead, false},
		{singleNode, false},
		{n1, false},
		{cycle1, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := hasCycle(tt.node); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCycleLen(t *testing.T) {
	var nilHead *single.ListNode

	singleNode := n(1)

	n1, n2, n3, n4, n5, n6, n7 := n(1), n(1), n(1), n(1), n(1), n(1), n(1)
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7

	cycle1, cycle2, cycle3, cycle4, cycle5, cycle6 := n(1), n(1), n(1), n(1), n(1), n(1)
	cycle1.Next = cycle2
	cycle2.Next = cycle3
	cycle3.Next = cycle4
	cycle4.Next = cycle5
	cycle5.Next = cycle6
	cycle6.Next = cycle3

	tests := []struct {
		node *single.ListNode
		want int
	}{
		{nilHead, 0},
		{singleNode, 0},
		{n1, 0},
		{cycle1, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findCycleLen(tt.node); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *single.ListNode {
	return &single.ListNode{Val: val}
}
