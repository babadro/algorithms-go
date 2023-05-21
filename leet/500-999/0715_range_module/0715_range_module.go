package _0715_range_module

import (
	"sort"
)

type interval struct {
	left, right int
}

type RangeModule struct {
	intervals []interval
}

func Constructor() RangeModule {
	return RangeModule{}
}

func (this *RangeModule) AddRange(left, right int) {
	idx := this.search(left)

	if idx == len(this.intervals) || this.intervals[idx].left > right {
		this.insert(idx, interval{left, right})
		return
	}

	i := idx
	for ; i < len(this.intervals); i++ {
		in := this.intervals[i]

		if in.left > right {
			break
		}

		left, right = min(in.left, left), max(in.right, right)
	}

	this.intervals[idx] = interval{left, right}

	this.shiftLeft(i, i-idx-1)
}

func (this *RangeModule) QueryRange(left, right int) bool {
	idx := this.search(left)

	if idx == len(this.intervals) {
		return false
	}

	in := this.intervals[idx]
	if left < in.left || right > in.right {
		return false
	}

	return true
}

func (this *RangeModule) RemoveRange(left, right int) {
	idx := this.search(left)

	if idx == len(this.intervals) || this.intervals[idx].left >= right {
		return
	}

	in := this.intervals[idx]

	if left > in.left {
		this.intervals[idx].right = left

		if right < in.right {
			this.insert(idx+1, interval{right, in.right})
			return
		}

		idx++
	}

	i := idx
	for ; i < len(this.intervals); i++ {
		in = this.intervals[i]

		if right < in.right {
			this.intervals[i].left = right

			break
		}
	}

	this.shiftLeft(i, i-idx)
}

func (this *RangeModule) insert(idx int, in interval) {
	this.intervals = append(this.intervals, interval{})

	copy(this.intervals[idx+1:], this.intervals[idx:len(this.intervals)-1])

	this.intervals[idx] = in
}

func (this *RangeModule) shiftLeft(startIDx, shift int) {
	if startIDx == len(this.intervals) || shift == 0 {
		return
	}

	copy(this.intervals[startIDx-shift:len(this.intervals)-shift], this.intervals[startIDx:])

	this.intervals = this.intervals[:len(this.intervals)-shift]
}

// find left-most interval that has right >= num
func (this *RangeModule) search(num int) int {
	return sort.Search(len(this.intervals), func(i int) bool {
		return this.intervals[i].right >= num
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
