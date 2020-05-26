package _056_merge_intervals

import "testing"

func TestMerge(t *testing.T) {
	t.Log(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	t.Log(merge([][]int{{1, 4}, {4, 5}}))
	t.Log(merge([][]int{{1, 4}, {5, 6}}))
	t.Log(merge([][]int{{1, 4}, {2, 3}}))
}
