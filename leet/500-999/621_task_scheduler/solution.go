package _621_task_scheduler

// https://leetcode.com/problems/task-scheduler/discuss/322606/simple-go-solution
// https://leetcode.com/problems/task-scheduler/discuss/341568/C%2B%2B-8ms-100-O(n)-solution-w-explanation
// passed. tptl medium (hard for me). #array
func leastInterval2(tasks []byte, n int) int {
	stat := [26]int{}
	for _, ch := range tasks {
		stat[ch-'A']++
	}

	maxFreq, maxFreqCount := 0, 0
	for _, freq := range stat {
		if maxFreq == freq {
			maxFreqCount++
		} else if maxFreq < freq {
			maxFreqCount = 1
			maxFreq = freq
		}
	}

	total := (maxFreq-1)*(n+1) + maxFreqCount

	return max(total, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
