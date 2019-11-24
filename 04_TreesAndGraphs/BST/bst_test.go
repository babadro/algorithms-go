package bst_test

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/BST"
	"testing"
)

func TestInorder(t *testing.T) {
	tree := makeTree()

	list := make([]int, 0)
	bst.Inorder(tree.Root, &list)
	for i := 0; i <= 10; i++ {
		if i != list[i] {
			t.Fatalf("Wrong order result Inorder func. want %v, got %v", i, list[i])
		}
	}
}

func TestSearch(t *testing.T) {
	tree := makeTree()

	node := bst.Search(tree.Root, 0)
	if node == nil {
		t.Fatal("Node must not be nil")
	}

	if node.Key != 0 {
		t.Errorf("Wrong key. Want %v, got %v", 0, node.Key)
	}
	node = bst.Search(tree.Root, 1000)
	if node != nil {
		t.Errorf("Node must be nil for non exists key")
	}
}

func TestMin(t *testing.T) {
	tree := makeTree()

	min := bst.Min(tree.Root)
	if min == nil {
		t.Fatal("Node must be not nil")
	}
	if min.Key != 0 {
		t.Errorf("Want 0, got %v", min.Key)
	}
}

func TestMax(t *testing.T) {
	tree := makeTree()

	max := bst.Max(tree.Root)
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
		node := bst.Search(tree.Root, i)
		successor := bst.Successor(node)
		want := i + 1
		if successor == nil {
			t.Fatalf("Want %v, got nil", want)
		}
		if successor.Key != want {
			t.Errorf("Want %v, got %v", want, successor.Key)
		}
	}
	largestNode := bst.Search(tree.Root, 10)
	successor := bst.Successor(largestNode)
	if successor != nil {
		t.Errorf("Want nil, got %v", successor.Key)
	}
}

func TestTree_Delete(t *testing.T) {
	tree := makeTree()

	for i := 0; i <= 10; i++ {
		nodeToDelete := bst.Search(tree.Root, i)
		tree.Delete(nodeToDelete)
		nilNode := bst.Search(tree.Root, i)
		if nilNode != nil {
			t.Errorf("Want nil, got node with key %v", nilNode.Key)
		}
	}
}

func makeTree() *bst.Tree {
	tree := &bst.Tree{}
	tree.Insert(&bst.Node{Key: 5})
	tree.Insert(&bst.Node{Key: 4})
	tree.Insert(&bst.Node{Key: 6})
	tree.Insert(&bst.Node{Key: 3})
	tree.Insert(&bst.Node{Key: 7})
	tree.Insert(&bst.Node{Key: 2})
	tree.Insert(&bst.Node{Key: 8})
	tree.Insert(&bst.Node{Key: 1})
	tree.Insert(&bst.Node{Key: 9})
	tree.Insert(&bst.Node{Key: 0})
	tree.Insert(&bst.Node{Key: 10})

	return tree
}
