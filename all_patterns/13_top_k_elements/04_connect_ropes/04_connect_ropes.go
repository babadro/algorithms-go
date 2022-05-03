package _4_connect_ropes

import (
	"container/heap"

	"github.com/babadro/algorithms-go/utils"
)

func minimumCostToConnectRopes(ropeLengths []int) int {
	minHeap := utils.IntMinHeap(ropeLengths)
	heap.Init(&minHeap)

	sum := 0
	for len(minHeap) > 1 {
		rope1 := heap.Pop(&minHeap).(int)
		rope2 := heap.Pop(&minHeap).(int)

		newRope := rope1 + rope2

		sum += newRope

		heap.Push(&minHeap, newRope)
	}

	return sum
}
