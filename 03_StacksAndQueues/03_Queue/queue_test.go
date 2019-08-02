package queue_test

import (
	"algorithms-go/03_StacksAndQueues/03_Queue"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := queue.New(12)

	queue.head = 6
	queue.tail = 6

	queue.Enqueue(15)
	queue.Enqueue(16)
	queue.Enqueue(9)
	queue.Enqueue(8)
	queue.Enqueue(4)

	queue.Enqueue(17)
	queue.Enqueue(3)
	queue.Enqueue(5)

	dequeueRes, _ := queue.Dequeue()
	if dequeueRes != 15 {
		t.Errorf("Dequeue(). got %v, want %v", dequeueRes, 15)
	}
}
