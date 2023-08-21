package _444_sequence_reconstruction

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	m := make(map[int]bool)
	for _, seq := range sequences {
		for _, num := range seq {
			m[num] = true
		}
	}

	if len(nums) > len(m) {
		return false
	}

}
