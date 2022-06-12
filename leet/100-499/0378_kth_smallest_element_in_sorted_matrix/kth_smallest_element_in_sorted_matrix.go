package _378_kth_smallest_element_in_sorted_matrix

import "container/heap"

// tptl passed. heaps solution
func kthSmallest(matrix [][]int, k int) int {
	h := minHeap{matrix: matrix}
	for i := 0; i < len(matrix) && i < k; i++ {
		heap.Push(&h, element{y: i, x: 0})
	}

	numberCount, res := 0, 0
	for h.Len() > 0 {
		el := heap.Pop(&h).(element)
		res = matrix[el.y][el.x]
		numberCount++
		if numberCount == k {
			break
		}

		el.x++
		if len(matrix[0]) > el.x {
			heap.Push(&h, el)
		}
	}

	return res
}

type element struct {
	y, x int
}

type minHeap struct {
	elements []element
	matrix   [][]int
}

func (a minHeap) Len() int {
	return len(a.elements)
}
func (a minHeap) Less(i, j int) bool {
	yI, xI, yJ, xJ := a.elements[i].y, a.elements[i].x, a.elements[j].y, a.elements[j].x
	return a.matrix[yI][xI] < a.matrix[yJ][xJ]
}
func (a minHeap) Swap(i, j int) {
	a.elements[i], a.elements[j] = a.elements[j], a.elements[i]
}
func (a *minHeap) Push(v interface{}) {
	a.elements = append(a.elements, v.(element))
}
func (a *minHeap) Pop() interface{} {
	last := len(a.elements) - 1
	res := a.elements[last]
	a.elements = a.elements[:last]

	return res
}
