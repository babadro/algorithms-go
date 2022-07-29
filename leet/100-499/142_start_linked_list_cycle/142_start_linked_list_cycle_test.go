package _142_start_linked_list_cycle

import (
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
)

func Test_detectCycle(t *testing.T) {
	head1 := single.ArrToLinkedList([]int{1, 2, 3, 4, 5, 6})
	single.FindNode(head1, 6).Next = single.FindNode(head1, 3)

	head2 := single.ArrToLinkedList([]int{1, 2, 3, 4, 5, 6})
	single.FindNode(head2, 6).Next = single.FindNode(head2, 4)

	head3 := single.ArrToLinkedList([]int{1, 2, 3, 4, 5, 6})
	single.FindNode(head3, 6).Next = single.FindNode(head3, 1)

	head4 := single.ArrToLinkedList([]int{1, 2})
	single.FindNode(head4, 2).Next = single.FindNode(head4, 1)

	noCycle := single.ArrToLinkedList([]int{1, 2, 3, 4, 5, 6})

	tests := []struct {
		head *single.ListNode
		want *single.ListNode
	}{
		{head1, single.FindNode(head1, 3)},
		{head2, single.FindNode(head2, 4)},
		{head3, single.FindNode(head3, 1)},
		{head4, single.FindNode(head4, 1)},
		{noCycle, nil},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := detectCycle(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCycleStart() = %d, want %d", got.Val, tt.want.Val)
			}
		})
	}
}
