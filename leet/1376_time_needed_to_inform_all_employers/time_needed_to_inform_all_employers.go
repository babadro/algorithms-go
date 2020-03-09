package _1376_time_needed_to_inform_all_employers

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
		totalTime += maxTime
		level = nextLevel
	}
	return totalTime
}
