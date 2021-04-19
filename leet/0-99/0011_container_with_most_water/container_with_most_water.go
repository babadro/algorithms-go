package _011_container_with_most_water

// TODO 2 this is brute force. Implement better solution.
func maxArea(height []int) int {
	max, curr, maxCurr, l := 0, 0, 0, len(height)
	for i := 0; i < l; i++ {
		maxCurr = 0
		for j := len(height) - 1; j > i; j-- {
			if maxPossible := height[i] * (j - i); maxCurr >= maxPossible {
				break
			}
			curr = min(height[j], height[i]) * (j - i)
			if curr > max {
				max = curr
			}
			if curr > maxCurr {
				maxCurr = curr
			}
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
