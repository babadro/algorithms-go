package _1376_time_needed_to_inform_all_employers

import "fmt"

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	managerToEmployees := make(map[int][]int)
	for employeeId, managerId := range manager {
		managerToEmployees[managerId] = append(managerToEmployees[managerId], employeeId)
	}
	level := []int{headID}
	var nextLevel []int
	totalTime := informTime[headID]
	for {
		nextLevel = make([]int, 0)
		nextLevelExists := false
		maxTime := 0
		for _, managerId := range level {
			if employees, ok := managerToEmployees[managerId]; ok {
				nextLevelExists = true
				for _, employeeId := range employees {
					if timeNeeded := informTime[employeeId]; timeNeeded > maxTime {
						maxTime = timeNeeded
					}
				}
				nextLevel = append(nextLevel, employees...)
			}
		}
		if !nextLevelExists {
			break
		}
		fmt.Println(maxTime)
		totalTime += maxTime
		level = nextLevel
	}
	return totalTime
}

func numOfMinutes4(n int, headID int, managers []int, informTime []int) int {
	if n == 1 {
		return 0
	}
	times := make([]int, n)
	times[headID] = informTime[headID]
	var dfs func(int) int
	dfs = func(id int) int {
		if times[id] != 0 {
			return times[id]
		}
		times[id] = dfs(managers[id]) + informTime[id]
		return times[id]
	}
	var maxTime int
	for id := range managers {
		maxTime = max(maxTime, dfs(id))
	}
	return maxTime
}

// TODO 3 понять
func numOfMinutes3(n int, headID int, manager []int, informTime []int) int {
	graph := make(map[int][]int)
	for i, m := range manager {
		graph[m] = append(graph[m], i)
	}

	return dfs(graph, headID, informTime)
}

func dfs(graph map[int][]int, start int, informTime []int) int {
	maxTime := 0
	if len(graph[start]) == 0 {
		return informTime[start]
	}
	for _, v := range graph[start] {
		maxTime = max(maxTime, dfs(graph, v, informTime))
	}
	return maxTime + informTime[start]
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
