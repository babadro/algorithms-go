package bst

import (
	"reflect"
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

func Test_Insert(t *testing.T) {
	tests := []struct {
		name        string
		key         int
		treeBuilder func() *Tree
		want        []int
	}{
		{
			"1",
			40,
			func() *Tree {
				n10, n20, n30, n100, n500 := n(10), n(20), n(30), n(100), n(500)
				n100.Left, n100.Right = n20, n500
				n20.Parent, n500.Parent = n100, n100

				n20.Left, n20.Right = n10, n30

				return &Tree{Root: n100}
			},
			[]int{10, 20, 30, 40, 100, 500},
		},
		{
			"2",
			1,
			func() *Tree {
				return &Tree{}
			},
			[]int{1},
		},
		{
			"3",
			19,
			func() *Tree {
				n10, n20, n30, n100, n500 := n(10), n(20), n(30), n(100), n(500)
				n100.Left, n100.Right = n20, n500
				n20.Parent, n500.Parent = n100, n100

				n20.Left, n20.Right = n10, n30

				return &Tree{Root: n100}
			},
			[]int{10, 19, 20, 30, 100, 500},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var arr []int
			f := func(node *Node) {
				arr = append(arr, node.Key)
			}
			tree, node := tt.treeBuilder(), &Node{Key: tt.key}
			tree.Insert(node)
			Inorder(tree.Root, f)
			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("got = %v, want %v", arr, tt.want)
			}
		})
	}
}

func n(key int) *Node {
	return &Node{Key: key}
}

func Test_Insert2(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			"1",
			[]int{50, 30, 20, 40, 70, 60, 80},
			[]int{20, 30, 40, 50, 60, 70, 80},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var arr []int
			f := func(node *Node) {
				arr = append(arr, node.Key)
			}
			tree := &Tree{Root: nil}
			for _, key := range tt.input {
				node := &Node{Key: key}
				tree.Insert(node)
			}
			Inorder(tree.Root, f)
			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("got = %v, want %v", arr, tt.want)
			}
		})
	}
}
