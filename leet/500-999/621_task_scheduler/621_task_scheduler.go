package _621_task_scheduler

import (
	"container/heap"
	"fmt"
)

// doesn't work
func leastInterval(tasks []byte, n int) int {
	m := make(map[byte]int)
	for _, b := range tasks {
		m[b]++
	}

	h := &minHeap{now: 0}

	for char, count := range m {
		h.arr = append(h.arr, [3]int{0, count, int(char)})
	}

	heap.Init(h)

	for h.arr[0][1] > 0 {
		if h.arr[0][0] <= h.now {
			h.arr[0][0] = h.now + 1 + n
			h.arr[0][1]--

			fmt.Printf("%s\n", string(byte(h.arr[0][2])))
		} else {
			fmt.Println("idle")
		}

		h.now++
		heap.Init(h)
	}

	return h.now
}

type minHeap struct {
	arr [][3]int
	now int
}

func (h minHeap) Len() int { return len(h.arr) }
func (h minHeap) Less(i, j int) bool {
	ok1, ok2 := h.now-h.arr[i][0] >= 0, h.now-h.arr[j][0] >= 0
	if ok1 == ok2 {
		return h.arr[i][1] > h.arr[j][1]
	}

	return ok1
}
func (h minHeap) Swap(i, j int)       { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }
func (h *minHeap) Push(x interface{}) { h.arr = append(h.arr, x.([3]int)) }
func (h *minHeap) Pop() (v interface{}) {
	h.arr, v = (h.arr)[:len(h.arr)-1], (h.arr)[len(h.arr)-1]
	return
}
