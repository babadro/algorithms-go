package linkedlist

type Node struct {
	Key  int
	Prev *Node
	Next *Node
}

type LinkedList struct {
	Nil *Node
}

func NewLinkedList() *LinkedList {
	nilNode := &Node{}
	nilNode.Prev = nilNode
	nilNode.Next = nilNode

	return &LinkedList{Nil: nilNode}
}

func (l *LinkedList) Search(k int) *Node {
	x := l.Nil.Next
	for x != l.Nil && x.Key != k {
		x = x.Next
	}
	return x
}

func (l *LinkedList) Insert(x *Node) {
	x.Next = l.Nil.Next
	l.Nil.Next.Prev = x
	l.Nil.Next = x
	x.Prev = l.Nil
}

func (l *LinkedList) Delete(x *Node) {
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
}
