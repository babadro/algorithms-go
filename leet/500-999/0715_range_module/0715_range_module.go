package _0715_range_module

import (
	"sort"

	"github.com/babadro/algorithms-go/utils"
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
	for ; i < len(this.intervals); idx++ {
		in := this.intervals[i]

		if in.left > right {
			break
		}

		left, right = utils.Min(in.left, left), utils.Max(in.right, right)
	}

	this.intervals[idx] = interval{left, right}

	this.shiftLeft(i, idx+1)
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

	i := idx
	for ; i < len(this.intervals); i++ {
		in := this.intervals[i]

		if right < in.right {
			this.intervals[i].left = right
			return
		}
	}

	this.shiftLeft(i, idx)
}

func (this *RangeModule) insert(idx int, in interval) {

}

func (this *RangeModule) shiftLeft(from, to int) {}

// find left-most interval that has right >= num
func (this *RangeModule) search(num int) int {
	return sort.Search(len(this.intervals), func(i int) bool {
		return this.intervals[i].right >= num
	})
}
