package bst

import (
	"testing"
)

func TestInorder(t *testing.T) {
	tree := makeTree()

	list := make([]int, 0)
	f := func(node *Node) {
		list = append(list, node.Key)
	}
	Inorder(tree.Root, f)
	for i := 0; i <= 10; i++ {
		if i != list[i] {
			t.Fatalf("Wrong order result Inorder func. want %v, got %v", i, list[i])
		}
	}
}

func TestSearch(t *testing.T) {
	tree := makeTree()

	node := Search(tree.Root, 0)
	if node == nil {
		t.Fatal("Node must not be nil")
	}

	if node.Key != 0 {
		t.Errorf("Wrong key. Want %v, got %v", 0, node.Key)
	}
	node = Search(tree.Root, 1000)
	if node != nil {
		t.Errorf("Node must be nil for non exists key")
	}
}

func TestMin(t *testing.T) {
	tree := makeTree()

	min := Min(tree.Root)
	if min == nil {
		t.Fatal("Node must be not nil")
	}
	if min.Key != 0 {
		t.Errorf("Want 0, got %v", min.Key)
	}
}

func TestMax(t *testing.T) {
	tree := makeTree()

	max := Max(tree.Root)
	if max == nil {
		t.Fatal("Node must be not nil")
	}
	if max.Key != 10 {
		t.Errorf("Want 10, got %v", max.Key)
	}
}

func TestSuccessor(t *testing.T) {
	tree := makeTree()

	for i := 0; i < 10; i++ {
		node := Search(tree.Root, i)
		successor := Successor(node)
		want := i + 1
		if successor == nil {
			t.Fatalf("Want %v, got nil", want)
		}
		if successor.Key != want {
			t.Errorf("Want %v, got %v", want, successor.Key)
		}
	}
	largestNode := Search(tree.Root, 10)
	successor := Successor(largestNode)
	if successor != nil {
		t.Errorf("Want nil, got %v", successor.Key)
	}
}

func TestTree_Delete(t *testing.T) {
	tree := makeTree()

	for i := 0; i <= 10; i++ {
		nodeToDelete := Search(tree.Root, i)
		tree.Delete(nodeToDelete)
		nilNode := Search(tree.Root, i)
		if nilNode != nil {
			t.Errorf("Want nil, got node with key %v", nilNode.Key)
		}
	}
}

func makeTree() *Tree {
	tree := &Tree{}
	tree.Insert(&Node{Key: 5})
	tree.Insert(&Node{Key: 4})
	tree.Insert(&Node{Key: 6})
	tree.Insert(&Node{Key: 3})
	tree.Insert(&Node{Key: 7})
	tree.Insert(&Node{Key: 2})
	tree.Insert(&Node{Key: 8})
	tree.Insert(&Node{Key: 1})
	tree.Insert(&Node{Key: 9})
	tree.Insert(&Node{Key: 0})
	tree.Insert(&Node{Key: 10})

	return tree
}
