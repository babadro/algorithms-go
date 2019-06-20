package linkedlist

type Node struct {
	Key  int
	prev *Node
	next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Search(k int) *Node {
	x := l.Head
	for x != nil && x.Key != k {
		x = x.next
	}
	return x
}

func (l *LinkedList) Insert(x *Node) {
	x.next = l.Head
	if l.Head != nil {
		l.Head.prev = x
	}
	l.Head = x
	x.prev = nil
}

func (l *LinkedList) Delete(x *Node) {
	if x.prev != nil {
		x.prev.next = x.next
	} else {
		l.Head = x.next
	}
	if x.next != nil {
		x.next.prev = x.prev
	}
}
