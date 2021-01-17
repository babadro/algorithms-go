package heap

// todo 1
type MaxHeap struct {
	list []int
}

func (m *MaxHeap) Build(source []int) {
	m.list = source
	for i := m.Size() / 2; i >= 0; i-- {
		m.Heapify(i)
	}
}

func (m *MaxHeap) Size() int {
	return len(m.list)
}

func (m *MaxHeap) Add(value int) {
	m.list = append(m.list, value)
	i := m.Size() - 1
	parent := Parent(i)

	for i > 0 && m.list[parent] < m.list[i] {
		m.list[i], m.list[parent] = m.list[parent], m.list[i]

		i = parent
		parent = Parent(parent)
	}
}

func (m *MaxHeap) Heapify(i int) {
	var left, right, largest int

	for {
		left = Left(i)
		right = Right(i)
		largest = i

		if left < m.Size() && m.list[left] > m.list[largest] {
			largest = left
		}

		if right < m.Size() && m.list[right] > m.list[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		m.list[i], m.list[largest] = m.list[largest], m.list[i]

		i = largest
	}
}

func (m *MaxHeap) ExtractMax() (int, bool) {
	n := len(m.list)
	if n == 0 {
		return 0, false
	}

	if n == 1 {
		res := m.list[0]
		m.list = m.list[:0]
		return res, true
	}

	root := m.list[0]

	m.list[0] = m.list[n-1]
	m.Heapify(0)

	return root, true
}

func Parent(idx int) int {
	return (idx - 1) / 2
}

func Left(idx int) int {
	return 2*idx + 1
}

func Right(idx int) int {
	return 2*idx + 2
}
