package _328_odd_even_linked_list

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

// 87% 100%
func oddEvenList(head *singly.ListNode) *singly.ListNode {
	if head == nil {
		return nil
	}

	oddCurr := head
	evenStart, evenCurr := head.Next, head.Next
	for evenCurr != nil && evenCurr.Next != nil {
		oddCurr.Next = oddCurr.Next.Next
		oddCurr = oddCurr.Next

		evenCurr.Next = evenCurr.Next.Next
		evenCurr = evenCurr.Next
	}

	oddCurr.Next = evenStart

	return head
}
