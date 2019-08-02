package queue

type Queue struct {
	array  []int
	head   int
	tail   int
	length int
}

func New(capacity int) *Queue {
	return &Queue{array: make([]int, capacity)}
}

func (q *Queue) Enqueue(x int) bool {
	if q.length == len(q.array) {
		return false
	}

	q.array[q.tail] = x
	if q.tail == len(q.array)-1 {
		q.tail = 0
	} else {
		q.tail++
	}

	q.length++
	return true
}

func (q *Queue) Dequeue() (int, bool) {
	if q.length == 0 {
		return 0, false
	}

	x := q.array[q.head]
	if q.head != q.tail {
		if q.head == len(q.array)-1 {
			q.head = 1
		} else {
			q.head++
		}
	}

	q.length--
	return x, true
}

func (q *Queue) Len() int {
	return q.length
}
