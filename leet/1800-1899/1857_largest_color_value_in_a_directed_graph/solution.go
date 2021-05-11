package _1857_largest_color_value_in_a_directed_graph

import "github.com/babadro/algorithms-go/utils"

// passed. hard. tptl
// todo 1 understand and outline
// todo 2 find faster or/and shorter solution like this: https://leetcode.com/problems/largest-color-value-in-a-directed-graph/discuss/1198639/because-only-26-colors-C%2B%2B-O(n)
// or this: https://leetcode.com/problems/largest-color-value-in-a-directed-graph/discuss/1200575/Topological-Sort-vs.-DFS-vs.-Dijkstra
func largestPathValue2(colors string, edges [][]int) int {
	vertexCount := len(colors)
	graph := make([][]int, vertexCount)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	max, hasCycle := -1, false
	memo := make([]map[byte]int, vertexCount)
	visited := make([]bool, vertexCount)

	for i := 0; i < vertexCount; i++ {
		dfs(i, graph, colors, visited, &hasCycle, memo, &max)

		if hasCycle {
			return -1
		}
	}

	return max
}

func dfs(curr int, graph [][]int, colors string, visited []bool, hasCycle *bool, memo []map[byte]int, max *int) map[byte]int {
	if visited[curr] {
		*hasCycle = true
		return nil
	}

	if dict := memo[curr]; dict != nil {
		return dict
	}

	visited[curr] = true
	currMap := make(map[byte]int)

	for _, vertex := range graph[curr] {
		resMap := dfs(vertex, graph, colors, visited, hasCycle, memo, max)

		if *hasCycle {
			return nil
		}

		for color, count := range resMap {
			currCount := currMap[color]
			currCount = utils.Max(currCount, count)
			currMap[color] = currCount
			*max = utils.Max(currCount, *max)
		}
	}

	currentNodeColorCount := currMap[colors[curr]]
	currMap[colors[curr]] = currentNodeColorCount + 1

	*max = utils.Max(currentNodeColorCount+1, *max)

	visited[curr] = false
	memo[curr] = currMap

	return currMap
}
