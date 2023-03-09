package _002_add_two_numbers

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

func addTwoNumbers(l1 *singly.ListNode, l2 *singly.ListNode) *singly.ListNode {
	var head, prev *singly.ListNode
	node1, node2 := l1, l2
	carry := false
	for node1 != nil || node2 != nil || carry {
		sum := 0
		if node1 != nil {
			sum += node1.Val
			node1 = node1.Next
		}

		if node2 != nil {
			sum += node2.Val
			node2 = node2.Next
		}

		if carry {
			sum += 1
		}

		if sum > 9 {
			sum %= 10
			carry = true
		} else {
			carry = false
		}

		if head == nil {
			prev = &singly.ListNode{Val: sum}
			head = prev
		} else {
			prev.Next = &singly.ListNode{Val: sum}
			prev = prev.Next
		}
	}

	return head
}
