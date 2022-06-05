package _234_palindrome_linked_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// todo 1 doesn't reassemble the order
func isPalindromeTheBest(head *single.ListNode) bool {
	slow, fast := head, head
	var prev *single.ListNode
	for fast != nil && fast.Next != nil {
		cur := slow
		slow, fast = slow.Next, fast.Next.Next
		// reverse first half of linked list
		cur.Next, prev = prev, cur
		// alternative implementation - todo 3 how does it work?
		//prev, slow.Next, slow, fast = slow, prev, slow.Next, fast.Next.Next
	}

	if fast != nil { // len(list) is odd
		slow = slow.Next
	}

	palindrome := true
	for prev != nil {
		if prev.Val != slow.Val {
			palindrome = false
		}

		prev, slow = prev.Next, slow.Next

		// reverse again to reassemble initial order
		// todo 1
	}

	return palindrome
}

// todo 1 doesn't reassemble last node.
func isPalindrome2(head *single.ListNode) bool {
	mid, end := head, head
	for end != nil && end.Next != nil {
		mid = mid.Next
		if end.Next.Next != nil {
			end = end.Next.Next
		} else {
			end = end.Next
			break
		}
	}

	// reverse
	var prev *single.ListNode
	cur := mid
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		cur, prev = next, cur
	}

	for left, right := head, end; right != nil; left, right = left.Next, right.Next {
		if left.Val != right.Val {
			return false
		}
	}

	return true
}
