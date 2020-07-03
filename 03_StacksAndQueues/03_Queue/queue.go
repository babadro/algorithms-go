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

func (q *Queue) Enqueue(x int) {
	if q.length == len(q.array) {
		var newCapacity int
		if len(q.array) == 0 {
			newCapacity = 1
		} else if len(q.array) < 1024 {
			newCapacity = len(q.array) * 2
		} else {
			newCapacity += len(q.array) / 4
		}
		newArr := make([]int, newCapacity)
		if q.tail == 0 {
			copy(newArr, q.array)
		} else {
			copy(newArr, q.array[q.head:])
			copy(newArr[len(q.array)-q.head:], q.array[:q.tail])
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
