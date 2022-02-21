package _2157_groups_of_strings

// tle, but correct
func groupStrings(words []string) []int {
	nums := wordsToNums(words)

	graph := make(map[int32][]int32)

	duplicates := make(map[int32]int)
	for _, num := range nums {
		if _, ok := graph[num]; ok {
			duplicates[num]++
		} else {
			graph[num] = nil
		}
	}

	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if connected(num, nums[j]) {
				graph[num] = append(graph[num], nums[j])
				graph[nums[j]] = append(graph[nums[j]], num)
			}
		}
	}

	maxGroupSize, groupsCount := 0, 0
	for len(graph) != 0 {
		groupsCount++
		randomStart := int32(0)
		for k := range graph {
			randomStart = k
			break
		}

		var duplicatesInGroup int
		visited := dfs(graph, randomStart, make(map[int32]bool), duplicates, &duplicatesInGroup)

		maxGroupSize = max(maxGroupSize, len(visited)+duplicatesInGroup)

		// clean graph
		for v := range visited {
			delete(graph, v)
		}
	}

	return []int{groupsCount, maxGroupSize}
}

func wordsToNums(words []string) []int32 {
	nums := make([]int32, len(words))
	for i, word := range words {
		var num int32
		for j := range word {
			bit := int32(1) << int32(word[j]-'a')
			num |= bit
		}

		nums[i] = num
	}

	return nums
}

func connected(a, b int32) bool {
	if a == b {
		return true
	}

	aBits, bBits := 0, 0
	for i := int32(0); i < 26; i++ {
		bit := int32(1) << i
		aBit, bBit := a&bit, b&bit
		if aBit != bBit {
			if aBit != 0 {
				aBits++
			} else {
				bBits++
			}

			if aBits == 2 || bBits == 2 {
				return false
			}
		}
	}

	return true
}

func dfs(g map[int32][]int32, v int32, visited map[int32]bool, duplicates map[int32]int, duplicatesRes *int) map[int32]bool {
	visited[v] = true
	*duplicatesRes += duplicates[v]
	for _, adj := range g[v] {
		if !visited[adj] {
			dfs(g, adj, visited, duplicates, duplicatesRes)
		}
	}

	return visited
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
