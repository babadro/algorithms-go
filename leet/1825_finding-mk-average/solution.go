package _1825_finding_mk_average

// todo 1 https://leetcode.com/problems/finding-mk-average/discuss/1155397/Java%3A-Binary-Search-(copied-from-Arrays-Class)-%2B-ArrayList-as-SlidingWindow
type MKAverage2 struct {
	m       int
	k       int
	resLen  int
	nums    []int
	indexes []int
}

func Constructor2(m int, k int) MKAverage2 {
	return MKAverage2{
		m:      m,
		k:      k,
		resLen: m - 2*k,
	}
}

func (mk *MKAverage2) slideByRemovingElement() {
	n, m := len(mk.nums), mk.m
	if n >= m {
		i := n - m
		copy(mk.indexes[i:len(mk.indexes)-1], mk.indexes[i+1:])
		mk.indexes = mk.indexes[:len(mk.indexes)-1]
	}
}

func (mk *MKAverage2) AddElement(num int) {
	mk.slideByRemovingElement()

	idx := mk.binarySearch(num)
	mk.indexes = append(mk.indexes, 0)
	copy(mk.indexes[idx+1:], mk.indexes[idx:len(mk.indexes)-1])
	mk.indexes[idx] = len(mk.nums)

	mk.nums = append(mk.nums, num)
}

func (mk *MKAverage2) CalculateMKAverage() int {
	if len(mk.nums) < mk.m {
		return -1
	}

	sum := 0
	for i := mk.m; i <= mk.m-mk.k-1; i++ {
		sum += mk.nums[mk.indexes[i]]
	}

	return sum / mk.resLen
}

func (mk *MKAverage2) binarySearch(key int) int {
	low, high := 0, len(mk.indexes)-1

	for low <= high {
		mid := (low + high) / 2
		midVal := mk.nums[mk.indexes[mid]]

		if midVal < key {
			low = mid + 1
		} else if midVal > key {
			high = mid - 1
		} else {
			return mid
		}
	}

	return low
}
