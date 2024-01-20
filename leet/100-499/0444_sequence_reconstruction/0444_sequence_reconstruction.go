package _444_sequence_reconstruction

// #bnsrg passed medium
// todo 2-3: check topological sort solution
func sequenceReconstruction(nums []int, sequences [][]int) bool {
	pairs := make(map[[2]int]bool)

	for _, seq := range sequences {
		for j := 1; j < len(seq); j++ {
			pairs[[2]int{seq[j-1], seq[j]}] = true
		}
	}

	for i := 1; i < len(nums); i++ {
		if !pairs[[2]int{nums[i-1], nums[i]}] {
			return false
		}
	}

	return true
}
