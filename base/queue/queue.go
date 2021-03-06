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

// todo 2 make unit tests for len > 1024
func (q *Queue) Enqueue(x int) {
	n := len(q.array)
	if q.length == n {
		var newCapacity int
		if n == 0 {
			newCapacity = 1
		} else if n < 1024 {
			newCapacity = n * 2
		} else {
			newCapacity = n + n/4
		}

		newArr := make([]int, newCapacity)
		if q.tail == 0 {
			copy(newArr, q.array)
		} else {
			copy(newArr, q.array[q.head:])
			copy(newArr[n-q.head:], q.array[:q.tail])
		}

		q.head = 0
		q.tail = q.length
		q.array = newArr
	}

	q.array[q.tail] = x
	if q.tail == len(q.array)-1 {
		q.tail = 0
	} else {
		q.tail++
	}

	q.length++

	return
}

func (q *Queue) Dequeue() (int, bool) {
	if q.length == 0 {
		return 0, false
	}

	x := q.array[q.head]
	if q.head == len(q.array)-1 {
		q.head = 0
	} else {
		q.head++
	}

	q.length--

	return x, true
}

func (q *Queue) Len() int {
	return q.length
}
