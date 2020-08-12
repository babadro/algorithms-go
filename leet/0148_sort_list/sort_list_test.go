package _148_sort_list

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSortList(t *testing.T) {
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	rnd := rand.New(src)
	t.Log(seed)

	for i := 0; i < 10; i++ {
		listLen := rnd.Intn(100)
		var head *single.ListNode
		var node *single.ListNode
		for j := 0; j < listLen; j++ {

			num := rand.Intn(20)
			if head == nil {
				head = &single.ListNode{Val: num}
				node = head
			} else {
				next := &single.ListNode{Val: num}
				node.Next = next
				node = node.Next
			}
		}
		arr := single.LinkedListToArr(head)
		newHead := sortList(head)
		sortedArr := single.LinkedListToArr(newHead)
		if len(arr) != len(sortedArr) {
			t.Errorf("arrayLen=%d, sortedArrayLen=%d", len(arr), len(sortedArr))
		}
		if !sort.IntsAreSorted(sortedArr) {
			t.Errorf("arr is not sorted\n: %v", sortedArr)
		}
	}
}
