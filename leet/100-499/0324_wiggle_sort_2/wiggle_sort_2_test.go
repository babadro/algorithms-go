package _324_wiggle_sort_2

import "testing"

func TestWiggleSort(t *testing.T) {
	nums1 := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums1)
	t.Log(nums1)

	nums2 := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums2)
	t.Log(nums2)

	nums3 := []int{1, 1, 1, 2, 2, 2, 3, 3}
	wiggleSort(nums3)
	t.Log(nums3)

	nums4 := []int{1, 1, 2, 1, 2, 2, 1}
	wiggleSort(nums4)
	t.Log(nums4)

	nums5 := []int{4, 5, 5, 6}
	wiggleSort(nums5)
	t.Log(nums5)
}
