package _142_start_linked_list_cycle

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// tptl. bes solution. floyd algorithm
func detectCycle(head *single.ListNode) *single.ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			break
		}
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

// tptl passed. medium
func detectCycle2(head *single.ListNode) *single.ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			break
		}
	}

	cycleLen := 0
	for {
		slow = slow.Next
		cycleLen++

		if slow == fast {
			break
		}
	}

	pointer1, pointer2 := head, head
	for i := 0; i < cycleLen; i++ {
		pointer2 = pointer2.Next
	}

	for pointer1 != pointer2 {
		pointer1 = pointer1.Next
		pointer2 = pointer2.Next
	}

	return pointer1
}