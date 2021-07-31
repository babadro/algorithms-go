package _1865_finding_pairs_with_a_certain_sum

// correct, but TLE
type FindSumPairs2 struct {
	sumToCount map[int]int
	sums       [][][2]int
}

func Constructor2(nums1 []int, nums2 []int) FindSumPairs2 {
	res := FindSumPairs2{
		sumToCount: make(map[int]int),
		sums:       make([][][2]int, len(nums2)),
	}

	for i, num2 := range nums2 {
		m := make(map[int]int, len(nums1))
		for _, num1 := range nums1 {
			sum := num1 + num2

			m[sum]++
			res.sumToCount[sum]++
		}

		for sum, count := range m {
			res.sums[i] = append(res.sums[i], [2]int{count, sum})
		}
	}

	return res
}

func (this *FindSumPairs2) Add(index int, val int) {
	sums := this.sums[index]
	for i := range sums {
		count, oldSum := sums[i][0], sums[i][1]
		newSum := oldSum + val
		this.sumToCount[oldSum] -= count
		this.sumToCount[newSum] += count

		sums[i][1] = newSum
	}
}

func (this *FindSumPairs2) Count(tot int) int {
	return this.sumToCount[tot]
}
