package queue

type Queue struct {
	array [12]int
	Head  int
	Tail  int
}

func (q *Queue) Enqueue(x int) {
	q.array[q.Tail] = x
	if q.Tail == len(q.array)-1 {
		q.Tail = 0
	} else {
		q.Tail++
	}
}

func (q *Queue) Dequeue() int {
	x := q.array[q.Head]
	if q.Head == len(q.array)-1 {
		q.Head = 1
	} else {
		q.Head++
	}
	return x
}
