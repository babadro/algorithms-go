package _703_kth_largest_element_in_a_stream

import (
	"container/heap"

	"github.com/babadro/algorithms-go/utils"
)

// tptl passed. easy
type KthLargest struct {
	minH utils.IntMinHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{
		minH: make(utils.IntMinHeap, 0, k+1),
		k:    k,
	}

	for _, num := range nums {
		kthLargest.Add(num)
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if len(this.minH) < this.k || this.minH[0] < val {
		heap.Push(&this.minH, val)
		if len(this.minH) > this.k {
			heap.Pop(&this.minH)
		}
	}

	return this.minH[0]
}
