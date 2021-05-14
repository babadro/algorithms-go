package _1825_finding_mk_average

import (
	"container/list"
	"github.com/babadro/algorithms-go/base/bst"
)

type MKAverage struct {
	queue  *list.List
	bst    bst.Tree
	k, m   int
	cached bool
	res    int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		queue: list.New(),
		k:     k,
		m:     m,
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

	sum := 0
	skip, counter := this.k, this.m-2*this.k

	f := func(node *bst.Node) {
		if skip > 0 {
			skip--
			return
		}

		if counter == 0 {
			return
		}

		sum += node.Key
		counter--
	}

	bst.Inorder(this.bst.Root, f)

	this.res = sum / this.queue.Len()
	this.cached = true

	return this.res
}
