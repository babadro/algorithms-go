package _070_climbing_stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	prev := 2
	prevPrev := 1
	for i := 2; i < n; i++ {
		sum := prev + prevPrev
		prevPrev = prev
		prev = sum
	}
	return prev
}
