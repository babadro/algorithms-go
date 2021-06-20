package contest

func minDifference(nums []int, queries [][]int) []int {
	return nil
}

func absDiff(a, b int) int {
	if res := a - b; res > 0 {
		return res
	} else if res < 0 {
		return -res
	} else {
		return -1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func absDiffArr(arr []int) int {
	return 0
}
