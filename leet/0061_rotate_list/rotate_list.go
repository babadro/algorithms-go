package _061_rotate_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// tptl. passed. todo 1 find for better solution (probably two pointers solution or solution
// without copying slices)
func rotateRight(head *single.ListNode, k int) *single.ListNode {
	nodes := make([]*single.ListNode, 0)

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

	tmp := make([]*single.ListNode, k)
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
