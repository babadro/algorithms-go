package _621_task_scheduler

import (
	"container/heap"
	"fmt"
)

// todo 1
func leastInterval(tasks []byte, n int) int {
	m := make(map[byte]int)
	for _, b := range tasks {
		m[b]++
	}

	h := &minHeap{}

	for char, count := range m {
		*h = append(*h, [3]int{0, count, int(char)})
	}

	heap.Init(h)

	curTime := 0
	for ; h.Len() > 0; curTime++ {
		if (*h)[0][0] > curTime {
			fmt.Println("idle")
			continue
		}

		item := heap.Pop(h).([3]int)
		item[0] = curTime + 1 + n
		item[1]--
		if item[1] > 0 {
			heap.Push(h, item)
		}

		fmt.Printf("%s\n", string(byte(item[2])))
	}

	return curTime
}

type minHeap [][3]int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	t1, t2 := h[i][0], h[j][0]
	if t1 == t2 {
		return h[i][1] > h[j][1]
	}

	return h[i][0] < h[j][0]
}
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([3]int)) }
func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}
