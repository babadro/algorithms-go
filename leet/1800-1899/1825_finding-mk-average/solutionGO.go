package _1825_finding_mk_average

// todo 1 need to understand
// passed. tptl. hard. best solution
type MKAverage struct {
	m, lk, mk, mn, p     int
	vals, cntBIT, valBIT []int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m:      m,
		lk:     k,
		mk:     m - k,
		mn:     m - k*2,
		p:      0,
		vals:   make([]int, 0, m),
		cntBIT: make([]int, 100_001),
		valBIT: make([]int, 100_001),
	}
}

func add(i, v int, bit []int) {
	for ; i < len(bit); i += i & (-i) { // todo 1 (i & (-i) lowest set bit
		bit[i] += v
	}
}

func getSum(i int, bit []int) (sum int) {
	for ; i > 0; i &= i - 1 { // todo 1
		sum += bit[i]
	}

	return sum
}

func (mk *MKAverage) AddElement(num int) {
	if len(mk.vals) == mk.m {
		del := mk.vals[mk.p]
		add(del, -del, mk.valBIT)
		add(del, -1, mk.cntBIT)

		mk.vals[mk.p] = num

		mk.p++
		if mk.p == mk.m {
			mk.p = 0
		}
	} else {
		mk.vals = append(mk.vals, num)
	}

	add(num, num, mk.valBIT)
	add(num, 1, mk.cntBIT)
}

func (mk *MKAverage) searchKSum(k int) int {
	l, r := 1, 100_001
	for l < r {
		m := (l + r) / 2
		if getSum(m, mk.cntBIT) >= k {
			r = m
		} else {
			l = m + 1
		}
	}

	return getSum(l, mk.valBIT) - (getSum(l, mk.cntBIT)-k)*l
}

func (mk *MKAverage) CalculateMKAverage() int {
	if len(mk.vals) < mk.m {
		return -1
	}

	return (mk.searchKSum(mk.mk) - mk.searchKSum(mk.lk)) / mk.mn
}
