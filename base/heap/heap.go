package heap

// todo 1 unit tests
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

// tptl
func (m *MaxHeap) Add(value int) {
	m.list = append(m.list, value)
	i := len(m.list) - 1
	parent := Parent(i)

	for i > 0 && m.list[parent] < m.list[i] {
		m.list[i], m.list[parent] = m.list[parent], m.list[i]

		i = parent
		parent = Parent(parent)
	}
}

func (m *MaxHeap) Delete(key int) {
	last := len(m.list) - 1
	m.list[key] = m.list[last]
	m.list = m.list[:last]

	// todo 1 fix for len(m.list) == 0
	m.Heapify(key)
}

// tptl
func (m *MaxHeap) Heapify(i int) {
	var left, right, largest int

	for {
		left = Left(i)
		right = Right(i)
		largest = i

		if left < len(m.list) && m.list[left] > m.list[largest] {
			largest = left
		}

		if right < m.Size() && m.list[right] > m.list[largest] {
			largest = right
		}

		if largest == i {
			return
		}

		m.list[i], m.list[largest] = m.list[largest], m.list[i]

		i = largest
	}
}

// tptl
func (m *MaxHeap) IncreaseKey(key, newVal int) {
	m.list[key] = newVal

	for key != 0 && m.list[key] < m.list[Parent(key)] {
		m.list[key], m.list[Parent(key)] = m.list[Parent(key)], m.list[key]

		key = Parent(key)
	}
}

// tptl
func (m *MaxHeap) DecreaseKey(key, newVal int) {
	m.list[key] = newVal
	m.Heapify(key)
}

// tptl
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

	last := n - 1
	m.list[0] = m.list[last]
	m.list = m.list[:last]

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
