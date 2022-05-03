package _973_k_closest_points_to_origin

import "container/heap"

// tptl. passed. medium
func kClosest(points [][]int, k int) [][]int {
	h := make(maxHeap, 0, k)
	for i := 0; i < k; i++ {
		heap.Push(&h, points[i])
	}

	for i := k; i < len(points); i++ {
		if squareDist(points[i]) < squareDist(h[0]) {
			heap.Pop(&h)
			heap.Push(&h, points[i])
		}
	}

	return h
}

type maxHeap [][]int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	return squareDist(h[i]) > squareDist(h[j])
}
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *maxHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}

func squareDist(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
