package _1376_time_needed_to_inform_all_employers

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	employeeMap := make(map[int][]int)
	for employeeId, managerId := range manager {
		employeeMap[managerId] = append(employeeMap[managerId], employeeId)
	}
	totalTime := informTime[headID]
	managerId := headID
	counter := 1
	for counter < n {
		employees := employeeMap[managerId]
		maxTime := 0
		for _, employeeId := range employees {
			if timeNeeded := informTime[employeeId]; timeNeeded > maxTime {
				maxTime = timeNeeded
			}
			counter++
		}
		totalTime += maxTime
	}
	return totalTime
}
