package _1865_finding_pairs_with_a_certain_sum

// tptl. passed. medium. #hashmap
type FindSumPairs struct {
	m1, m2 map[int]int
	nums2  []int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	m1, m2 := make(map[int]int), make(map[int]int)
	for _, num1 := range nums1 {
		m1[num1]++
	}

	for _, num2 := range nums2 {
		m2[num2]++
	}

	return FindSumPairs{
		m1:    m1,
		m2:    m2,
		nums2: nums2,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	oldNum2 := this.nums2[index]
	newNum2 := oldNum2 + val
	this.nums2[index] = newNum2

	this.m2[oldNum2]--
	if this.m2[oldNum2] == 0 {
		delete(this.m2, oldNum2)
	}

	this.m2[newNum2]++
}

func (this *FindSumPairs) Count(tot int) int {
	var res int
	for num1, count1 := range this.m1 {
		num2 := tot - num1
		if count2, ok := this.m2[num2]; ok {
			res += count1 * count2
		}
	}

	return res
}
