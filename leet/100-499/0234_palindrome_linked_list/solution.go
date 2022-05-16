package _234_palindrome_linked_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

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
