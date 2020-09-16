package _328_odd_even_linked_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// 87% 100%
func oddEvenList(head *single.ListNode) *single.ListNode {
	if head == nil {
		return nil
	}

	oddStart, oddCurr := head, head
	evenStart, evenCurr := oddStart.Next, oddStart.Next
	for evenCurr != nil && evenCurr.Next != nil {
		oddCurr.Next = oddCurr.Next.Next
		oddCurr = oddCurr.Next

		evenCurr.Next = evenCurr.Next.Next
		evenCurr = evenCurr.Next
	}

	oddCurr.Next = evenStart

	return oddStart
}
