package _2179_count_good_triplets_in_array

const n = 100000

var bt = make([]int, n+1)

// todo 1 doesn't work
func goodTriplets(nums1 []int, nums2 []int) int64 {
	res, sz := 0, len(nums1)
	ids := make([]int, sz)

	for i := 0; i < sz-1; i++ {
		ids[nums2[i]] = i
	}

	for i, num1 := range nums1 {
		mid := ids[num1]
		sm := prefixSum(mid)
		gr := sz - 1 - mid - (i - sm)
		res += sm * gr
		add(mid, 1)
	}

	return int64(res)
}

func prefixSum(i int) int {
	sum := 0
	for i = i + 1; i > 0; i -= i & (-i) {
		sum += bt[i]
	}

	return sum
}

func add(i, val int) {
	for i = i + 1; i <= n; i += i & (-i) {
		bt[i] += val
	}
}
