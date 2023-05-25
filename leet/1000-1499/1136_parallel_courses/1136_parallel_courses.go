package _1136_parallel_courses

// #sgbrn
// You are given an integer n, which indicates that there are n courses labeled from 1 to n.
// You are also given an array relations where relations[i] = [prevCoursei, nextCoursei], representing a prerequisite relationship between course prevCoursei and course nextCoursei: course prevCoursei has to be taken before course nextCoursei.
// In one semester, you can take any number of courses as long as you have taken all the prerequisites in the previous semester for the courses you are taking.
// Return the minimum number of semesters needed to take all courses. If there is no way to take all the courses, return -1.
func minimumSemesters(n int, relations [][]int) int {
	g, inDegree := make(map[int][]int), make(map[int]int)

	for _, r := range relations {
		src, dst := r[0], r[1]
		g[src] = append(g[src], dst)
		inDegree[dst]++
	}

	var sources []int
	for v := 1; v <= n; v++ {
		if inDegree[v] == 0 {
			sources = append(sources, v)
		}
	}

	semestersNum := 0
	for i := 0; i < len(sources); {
		coursesInSemester := len(sources) - i
		semestersNum++
		for j := 0; j < coursesInSemester; j++ {
			src := sources[j+i]
			for _, dst := range g[src] {
				inDegree[dst]--
				if inDegree[dst] == 0 {
					sources = append(sources, dst)
				}
			}
		}

		i += coursesInSemester
	}

	if len(sources) < n {
		return -1
	}

	return semestersNum
}
