package _2157_groups_of_strings

// tptl. passed. best solution
// https://leetcode.com/problems/groups-of-strings/discuss/1731073/Union-Find-vs.-DFS
func groupStrings2(words []string) []int {
	m := make(map[int32]int16)
	for _, w := range words {
		var mask int32
		for _, ch := range w {
			mask |= 1 << (ch - 'a')
		}

		m[mask]++
	}

	groupsCount, maxGroupSize := 0, 0
	for len(m) > 0 {
		var anyMask int32 // peek first random mask
		for mask := range m {
			anyMask = mask
			break
		}

		size := dfs2(m, anyMask)

		if size > 0 {
			maxGroupSize = max(maxGroupSize, size)
			groupsCount++
		}
	}

	return []int{groupsCount, maxGroupSize}
}

func dfs2(m map[int32]int16, mask int32) int {
	res := 0
	if count, ok := m[mask]; ok {
		res += int(count)
		delete(m, mask)

		for i := 0; i < 26; i++ {
			res += dfs2(m, mask^(1<<i))
			for j := i + 1; j < 26; j++ {
				if (mask>>i)&1 != (mask>>j)&1 {
					res += dfs2(m, mask^(1<<i)^(1<<j))
				}
			}
		}
	}

	return res
}
