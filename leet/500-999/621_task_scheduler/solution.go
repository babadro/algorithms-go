package _621_task_scheduler

// https://leetcode.com/problems/task-scheduler/discuss/322606/simple-go-solution
https://leetcode.com/problems/task-scheduler/discuss/341568/C%2B%2B-8ms-100-O(n)-solution-w-explanation

func leastInterval2(tasks []byte, n int) int {
	count := [26]int{}
	for _, b := range tasks {
		count[b-'A']++
	}

	maxCount := 0
	for _, c := range count {
		maxCount = max(c, maxCount)
	}

	return max()


}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
