package _6_reconstructing_a_sequence

// see premium leetcode https://leetcode.com/problems/sequence-reconstruction/
func canConstruct(originalSeq []int, sequences [][]int) bool {
	g, inDegree := make(map[int][]int), make(map[int]int)

	for _, seq := range sequences {
		for i, src := range seq {
			if _, ok := inDegree[src]; !ok {
				inDegree[src] = 0
			}

			if i == len(seq)-1 {
				break
			}

			dst := seq[i+1]
			g[src] = append(g[src], dst)
			inDegree[dst]++
		}
	}

	// if we don't have ordering rules for all the numbers
	// we'll not able to  uniquely construct the sequence
	if len(inDegree) != len(originalSeq) {
		return false
	}

	var sources []int
	for v, degree := range inDegree {
		if degree == 0 {
			sources = append(sources, v)
		}
	}

	// todo 3 can be optimized, we don't really need to store the whole sources arr here
	for i := 0; i < len(sources); i++ {
		// more than one remaining sources mean, there is more than one way to reconstruct the sequence
		if len(sources)-i > 1 {
			return false
		}

		// the next source (or number) is different from the original sequence
		if originalSeq[i] != sources[i] {
			return false
		}

		src := sources[i]
		for _, dst := range g[src] {
			inDegree[dst]--
			if inDegree[dst] == 0 {
				sources = append(sources, dst)
			}
		}
	}

	return len(sources) == len(originalSeq)
}
