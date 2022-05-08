package _23_merge_k_sorted_lists

import (
	"container/heap"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
)

// tptl. passed.
func mergeKLists(lists []*single.ListNode) *single.ListNode {
	h := make(listItems, 0, len(lists))
	for listIDx, li := range lists {
		if li != nil {
			h = append(h, [2]int{li.Val, listIDx})
			lists[listIDx] = li.Next
		}
	}

	if len(h) == 0 {
		return nil
	}
	heap.Init(&h)

	var head, prev *single.ListNode

	for len(h) > 0 {
		minItem := heap.Pop(&h).([2]int)
		val, listIDx := minItem[0], minItem[1]

		newNode := &single.ListNode{Val: val}

		if head == nil {
			head, prev = newNode, newNode
		} else {
			prev.Next = newNode
			prev = newNode
		}

		if li := lists[listIDx]; li != nil {
			heap.Push(&h, [2]int{li.Val, listIDx})
			lists[listIDx] = li.Next
		}
	}

	return head
}

type listItems [][2]int

func (li listItems) Len() int {
	return len(li)
}
func (li listItems) Less(i, j int) bool {
	return li[i][0] < li[j][0]
}
func (li listItems) Swap(i, j int) {
	li[i], li[j] = li[j], li[i]
}
func (li *listItems) Push(v interface{}) {
	*li = append(*li, v.([2]int))
}
func (li *listItems) Pop() interface{} {
	last := len(*li) - 1
	res := (*li)[last]
	*li = (*li)[:last]

	return res
}
