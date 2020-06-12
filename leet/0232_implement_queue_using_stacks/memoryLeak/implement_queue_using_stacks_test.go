package memoryLeak

import "testing"

func TestMyQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	t.Log(queue.Peek())
	t.Log(queue.Empty())
	t.Log(queue.Pop())
	t.Log(queue.Pop())
	t.Log(queue.Pop())
	t.Log(queue.Empty())
}
