package _142_start_linked_list_cycle

import (
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

func Test_detectCycle(t *testing.T) {
	head1 := singly.ArrToLinkedList([]int{1, 2, 3, 4, 5, 6})
	singly.FindNode(head1, 6).Next = singly.FindNode(head1, 3)

	head2 := singly.ArrToLinkedList([]int{1, 2, 3, 4, 5, 6})
	singly.FindNode(head2, 6).Next = singly.FindNode(head2, 4)

	head3 := singly.ArrToLinkedList([]int{1, 2, 3, 4, 5, 6})
	singly.FindNode(head3, 6).Next = singly.FindNode(head3, 1)

	head4 := singly.ArrToLinkedList([]int{1, 2})
	singly.FindNode(head4, 2).Next = singly.FindNode(head4, 1)

	noCycle := singly.ArrToLinkedList([]int{1, 2, 3, 4, 5, 6})

	tests := []struct {
		head *singly.ListNode
		want *singly.ListNode
	}{
		{head1, singly.FindNode(head1, 3)},
		{head2, singly.FindNode(head2, 4)},
		{head3, singly.FindNode(head3, 1)},
		{head4, singly.FindNode(head4, 1)},
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
