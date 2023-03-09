package _061_rotate_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"

// tptl. passed. todo2 look two pointer solution
func rotateRight(head *singly.ListNode, k int) *singly.ListNode {
	if head == nil {
		return nil
	}

	length := 1
	oldTail := head
	for ; oldTail.Next != nil; oldTail, length = oldTail.Next, length+1 {
	}

	k %= length
	if k == 0 {
		return head
	}

	newTail := head
	for i := 0; i < length-1-k; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next

	newTail.Next = nil
	oldTail.Next = head

	return newHead
}

// passed. naive
func rotateRight2(head *singly.ListNode, k int) *singly.ListNode {
	nodes := make([]*singly.ListNode, 0)

	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	n := len(nodes)
	if n == 0 {
		return head
	}

	k %= n

	if k == 0 {
		return head
	}

	tmp := make([]*singly.ListNode, k)
	copy(tmp, nodes[n-k:])
	copy(nodes[k:], nodes[:n-k])
	copy(nodes[:k], tmp)

	for i := range nodes {
		if i == n-1 {
			nodes[i].Next = nil
		} else {
			nodes[i].Next = nodes[i+1]
		}
	}

	return nodes[0]
}
