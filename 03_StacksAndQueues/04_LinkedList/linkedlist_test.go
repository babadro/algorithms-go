package linkedlist_test

import (
	"algorithms-go/03_StacksAndQueues/04_LinkedList"
	"testing"
)

func TestLinkedList(t *testing.T) {
	var l linkedlist.LinkedList

	l.Insert(&linkedlist.Node{Key: 1})
	l.Insert(&linkedlist.Node{Key: 4})
	l.Insert(&linkedlist.Node{Key: 16})
	l.Insert(&linkedlist.Node{Key: 9})
	l.Insert(&linkedlist.Node{Key: 25})

	searchRes := l.Search(4)
	if searchRes == nil {
		t.Errorf("search result must be non nil")
	}

	if searchRes.Key != 4 {
		t.Errorf("search got %v, want %v", searchRes.Key, 4)
	}

	head := l.Head
	if head == nil {
		t.Error("head must be non nil")
	}

	if head.Key != 25 {
		t.Errorf("head is %v, want %v", head.Key, 25)
	}

	nodeToDelete := l.Search(4)
	l.Delete(nodeToDelete)

	if l.Search(4) != nil {
		t.Error("Deleted node still exists")
	}
}
