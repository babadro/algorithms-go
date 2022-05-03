package _3_k_closest_points_to_the_origin

import "container/heap"

// leetcode 973
// Given an array of points in a 2D plane, find ‘K’ closest points to the origin.
func findClosestPoints(points [][2]int, k int) [][2]int {
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

type maxHeap [][2]int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	return squareDist(h[i]) > squareDist(h[j])
}
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(v interface{}) { *h = append(*h, v.([2]int)) }
func (h *maxHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}

func squareDist(point [2]int) int {
	return point[0]*point[0] + point[1]*point[1]
}
