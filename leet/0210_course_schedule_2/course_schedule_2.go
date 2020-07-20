package _210_course_schedule_2

// TODO 1 doesn't work
// input 3, [[0,1],[0,2],[1,2]]
// output [2, 0, 1]
// expected [2, 1, 0]
func findOrder1(numCourses int, prerequisites [][]int) []int {
	set := make(map[int]int, numCourses)
	for _, arr := range prerequisites {
		set[arr[0]] = arr[1]
	}
	var res []int
	done := make(map[int]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if _, ok := set[i]; !ok {
			res = append(res, i)
			done[i] = true
		}
	}
	if len(res) == 0 {
		return nil
	}

	for len(set) > 0 {
		k, deleteKey := 0, false
		for course, prerequisite := range set {
			if done[prerequisite] {
				done[course] = true
				res = append(res, course)
				k, deleteKey = course, true
				break
			}
		}
		if !deleteKey {
			return nil
		}
		delete(set, k)
	}

	if len(res) < numCourses {
		return nil
	}
	return res
}

/*
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for _, arr := range prerequisites {
		graph[arr[0]] = append(graph[arr[0]], arr[1])
	}

	counter := 0
	courses := make(map[int]bool)

	for i, arr := range graph {
		for _, j := range arr {
			visited := make(map[int]bool)
			if visited[j] {
				return nil
			}
			visited[j] = true

		}
	}

}

func add(v int, res *[]int, graph [][]int, visited, courses map[int]bool) {
	if len(graph[v]) == 0 {

	}
}*/
