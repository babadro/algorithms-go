package _1834_single_threaded_cpu

import (
	"container/heap"
	"sort"
)

// passed. tptl. medium (hard for me)
func getOrder(tasks [][]int) []int {
	n := len(tasks)
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}

	sort.Slice(indexes, func(i, j int) bool {
		return tasks[indexes[i]][0] < tasks[indexes[j]][0]
	})

	h := &minHeap{tasks: tasks}

	res := make([]int, 0, n)
	for i, t := 0, 1; len(res) < n; {
		for i < n && tasks[indexes[i]][0] <= t {
			heap.Push(h, indexes[i])
			i++
		}

		if h.Len() > 0 {
			idx := heap.Pop(h).(int)
			t += tasks[idx][1]

			res = append(res, idx)
		} else {
			t = tasks[indexes[i]][0]
		}
	}

	return res
}

type minHeap struct {
	indexes []int
	tasks   [][]int
}

func (m minHeap) Len() int { return len(m.indexes) }
func (m minHeap) Less(i, j int) bool {
	procTime1, procTime2 := m.tasks[m.indexes[i]][1], m.tasks[m.indexes[j]][1]
	if procTime1 != procTime2 {
		return procTime1 < procTime2
	}

	return m.indexes[i] < m.indexes[j]
}
func (m minHeap) Swap(i, j int) { m.indexes[i], m.indexes[j] = m.indexes[j], m.indexes[i] }
func (m *minHeap) Push(v interface{}) {
	(*m).indexes = append((*m).indexes, v.(int))
}
func (m *minHeap) Pop() interface{} {
	last := len((*m).indexes) - 1
	res := (*m).indexes[last]
	(*m).indexes = (*m).indexes[:last]

	return res
}
