package _1865_finding_pairs_with_a_certain_sum


// todo 1
type FindSumPairs struct {
	sumToCount map[int]int
	sums []map[int]int
	nums1, nums2 []int
}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	res := FindSumPairs{
		sumToCount: make(map[int]int),
		pairs:      make([][]pair, len(nums2)),
		nums1:      nums1,
		nums2:      nums2,
	}

	for i, num2 := range nums2 {
		for j, num1 := range nums1 {
			sum := num1 + num2

			res.
		}
	}


}


func (this *FindSumPairs) Add(index int, val int)  {

}


func (this *FindSumPairs) Count(tot int) int {

}

func (this *FindSumPairs) recalc() {

}