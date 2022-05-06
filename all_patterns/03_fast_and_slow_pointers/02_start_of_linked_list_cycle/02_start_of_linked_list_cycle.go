package _2_start_of_linked_list_cycle

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// leetcode 142
func findCycleStart(head *single.ListNode) *single.ListNode {
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
