package single

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeRandom struct {
	Val    int
	Next   *ListNodeRandom
	Random *ListNodeRandom
}

func ArrToLinkedList(arr []int) *ListNode {
	var node *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		newNode := &ListNode{Val: arr[i], Next: node}
		node = newNode
	}
	return node
}

func LinkedListToArr(head *ListNode) (arr []int) {
	node := head
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}
	return arr
}

func ArrToLinkedListRandom(arr [][2]int) *ListNodeRandom {
	if len(arr) == 0 {
		return nil
	}
	var nodesArr []*ListNodeRandom
	for i := 0; i < len(arr); i++ {
		node := &ListNodeRandom{}
		node.Val = arr[i][0]
		nodesArr = append(nodesArr, node)
	}
	for i := 0; i < len(arr); i++ {
		if randomIdx := arr[i][1]; randomIdx != -1 {
			nodesArr[i].Random = nodesArr[randomIdx]
		}
		if next := i + 1; next < len(arr) {
			nodesArr[i].Next = nodesArr[next]
		}
	}

	return nodesArr[0]
}

func LinkedListRandomToArr(head *ListNodeRandom) (arr [][2]int) {
	node := head
	for node != nil {
		randomVal := -1
		if node.Random != nil {
			randomVal = node.Random.Val
		}
		arr = append(arr, [2]int{node.Val, randomVal})
		node = node.Next
	}
	return arr
}

func FindNode(node *ListNode, val int) *ListNode {
	for ; node != nil; node = node.Next {
		if node.Val == val {
			return node
		}
	}

	return nil
}
