package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New(12)

	q.head = 6
	q.tail = 6

	q.Enqueue(15)
	q.Enqueue(16)
	q.Enqueue(9)
	q.Enqueue(8)
	q.Enqueue(4)

	q.Enqueue(17)
	q.Enqueue(3)
	q.Enqueue(5)

	dequeueRes, _ := q.Dequeue()
	if dequeueRes != 15 {
		t.Errorf("Dequeue(). got %v, want %v", dequeueRes, 15)
	}
}
func TestQueue2(t *testing.T) {
	q := New(3)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	dequeueRes, _ := q.Dequeue()
	t.Log(dequeueRes)
	dequeueRes, _ = q.Dequeue()
	t.Log(dequeueRes)
}

func TestQueue3(t *testing.T) {
	a := make([]int, 0) // TODO 1
	t.Log(a)
}
