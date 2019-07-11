package bst_test

import (
	"algorithms-go/04_TreesAndGraphs/BST"
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
