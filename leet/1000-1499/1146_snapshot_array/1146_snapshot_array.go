package _1146_snapshot_array

import "slices"

type entry struct {
	value, snapshotID int
}

type SnapshotArray struct {
	arr     [][]entry
	counter int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{arr: make([][]entry, length), counter: 0}
}

func (this *SnapshotArray) Set(index int, val int) {
	if len(this.arr[index]) == 0 {
		this.arr[index] = []entry{{value: val, snapshotID: this.counter}}

		return
	}

	last := len(this.arr[index]) - 1

	if this.arr[index][last].snapshotID == this.counter {
		this.arr[index][last].value = val

		return
	}

	this.arr[index] = append(this.arr[index], entry{value: val, snapshotID: this.counter})
}

func (this *SnapshotArray) Snap() int {
	this.counter++

	return this.counter - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	entries := this.arr[index]

	if len(entries) == 0 {
		return 0
	}

	compare := func(a entry, b int) int {
		return a.snapshotID - b
	}

	idx, found := slices.BinarySearchFunc(this.arr[index], snap_id, compare)
	if found {
		return entries[idx].value
	}

	if idx == 0 {
		return 0
	}

	return entries[idx-1].value
}
