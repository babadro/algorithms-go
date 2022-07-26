package _5_alien_dictionary

// see premium leetcode https://leetcode.com/problems/alien-dictionary/
func findOrder(words []string) string {
	g, inDegree := make(map[byte][]byte), make(map[byte]int)

	for i := 0; i < len(words)-1; i++ {
		j := 0
		for ; j < len(words[i]) && j < len(words[i+1]); j++ {
			if words[i][j] != words[i+1][j] {
				break
			}
		}

		if j == len(words[i]) || j == len(words[i+1]) {
			continue
		}

		src, dst := words[i][j], words[i+1][j]
		g[src] = append(g[src], dst)
		inDegree[dst]++
		if _, ok := inDegree[src]; !ok {
			inDegree[src] = 0
		}
	}

	var sources []byte
	for v, degree := range inDegree {
		if degree == 0 {
			sources = append(sources, v)
		}
	}

	for i := 0; i < len(sources); i++ {
		src := sources[i]
		for _, dst := range g[src] {
			inDegree[dst]--
			if inDegree[dst] == 0 {
				sources = append(sources, dst)
			}
		}
	}

	return string(sources)
}
