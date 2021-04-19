package _384_shuffle_an_array

import (
	"math/rand"
	"sort"
)

// 80% 92%
type Solution struct {
	arr  []int
	idxs []int
}

func (this *Solution) Len() int {
	return len(this.arr)
}

func (this *Solution) Less(i, j int) bool {
	return this.idxs[i] < this.idxs[j]
}

func (this *Solution) Swap(i, j int) {
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
	this.idxs[i], this.idxs[j] = this.idxs[j], this.idxs[i]
}

func Constructor(nums []int) Solution {
	n := len(nums)
	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	return Solution{
		arr:  nums,
		idxs: idxs,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	sort.Sort(this)

	return this.arr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := range this.arr {
		idx := rand.Intn(len(this.arr)-i) + i
		this.Swap(i, idx)
	}

	return this.arr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
