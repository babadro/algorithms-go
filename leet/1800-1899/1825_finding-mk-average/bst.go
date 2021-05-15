package _1825_finding_mk_average

import (
	"container/list"
	"github.com/babadro/algorithms-go/base/bst"
)

type MKAverage struct {
	queue        *list.List
	bst          bst.Tree
	m, k, length int
	cached       bool
	res          int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		queue:  list.New(),
		m:      m,
		k:      k,
		length: m - 2*k,
	}
}

func (this *MKAverage) AddElement(num int) {
	if this.queue.Len() == this.m {
		nodeToRemove := this.queue.Front()

		this.queue.Remove(nodeToRemove)
		this.bst.Delete(nodeToRemove.Value.(*bst.Node))
	}

	node := &bst.Node{
		Key: num,
	}

	this.bst.Insert(node)
	this.queue.PushBack(node)

	this.cached = false
}

func (this *MKAverage) CalculateMKAverage() int {
	if this.queue.Len() < this.m {
		return -1
	}

	if this.cached {
		return this.res
	}

	sum, skip, counter := 0, 0, 0

	f := func(node *bst.Node) {
		if skip < this.k {
			skip++
			return
		}

		if counter == this.length {
			return
		}

		sum += node.Key
		counter++
	}

	bst.Inorder(this.bst.Root, f)

	this.res = sum / this.length
	this.cached = true

	return this.res
}
