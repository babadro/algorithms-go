package _021_merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	head := new(ListNode)
	node := head
	prev := head
	for {
		if l1 == nil && l2 == nil {
			prev.Next = nil
			break
		}
		if l1 == nil || (l2 != nil && l1.Val > l2.Val) {
			node.Val = l2.Val
			l2 = l2.Next
		} else {
			node.Val = l1.Val
			l1 = l1.Next
		}
		node.Next = new(ListNode)
		prev = node
		node = node.Next
	}
	return head
}
