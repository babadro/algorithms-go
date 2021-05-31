package _1882_process_tasks_using_servers

import (
	"container/heap"
)

// passed. tptl. mine! medium.
// it's too slow todo 2 check another solution, it has two cycles only https://leetcode.com/problems/process-tasks-using-servers/discuss/1241936/Two-Heaps
func assignTasks(servers []int, tasks []int) []int {
	pool := &poolMinHeap{servers: servers}
	for i := 0; i < len(servers); i++ {
		heap.Push(pool, i)
	}

	busy := &busyMinHeap{}

	res := make([]int, 0, len(tasks))

	i := 0
	for currentTime := 0; busy.Len() > 0 || currentTime < len(tasks); currentTime++ {
		for busy.Len() > 0 && (*busy)[0].finish <= currentTime {
			finishedTask := heap.Pop(busy).(task)

			heap.Push(pool, finishedTask.serverIDx)
		}

		for ; pool.Len() > 0 && i <= currentTime && i < len(tasks); i++ {
			serverIDx := heap.Pop(pool).(int)

			res = append(res, serverIDx)

			newTask := task{finish: currentTime + tasks[i], serverIDx: serverIDx}

			heap.Push(busy, newTask)
		}
	}

	return res
}

type poolMinHeap struct {
	servers []int
	idxes   []int
}

func (h poolMinHeap) Len() int { return len(h.idxes) }
func (h poolMinHeap) Less(i, j int) bool {
	idxL, idxR := h.idxes[i], h.idxes[j]
	l, r := h.servers[idxL], h.servers[idxR]
	if l != r {
		return l < r
	}

	return idxL < idxR
}
func (h poolMinHeap) Swap(i, j int)       { h.idxes[i], h.idxes[j] = h.idxes[j], h.idxes[i] }
func (h *poolMinHeap) Push(v interface{}) { h.idxes = append(h.idxes, v.(int)) }
func (h *poolMinHeap) Pop() interface{} {
	last := len(h.idxes) - 1
	res := h.idxes[last]
	h.idxes = h.idxes[:last]

	return res
}

type task struct {
	finish    int
	serverIDx int
}

type busyMinHeap []task

func (h busyMinHeap) Len() int { return len(h) }
func (h busyMinHeap) Less(i, j int) bool {
	return h[i].finish < h[j].finish
}
func (h busyMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *busyMinHeap) Push(v interface{}) {
	*h = append(*h, v.(task))
}
func (h *busyMinHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
